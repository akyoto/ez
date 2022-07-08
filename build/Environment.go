package build

import (
	"strings"
	"sync"
	"sync/atomic"

	"github.com/akyoto/q/build/types"
)

// Environment represents the global state.
type Environment struct {
	Packages        map[string]*Package
	Functions       map[string]*Function
	Types           map[string]*types.Type
	StandardLibrary string
}

// NewEnvironment creates a new build environment.
func NewEnvironment() (*Environment, error) {
	standardLibrary, err := FindStandardLibrary()

	if err != nil {
		return nil, err
	}

	environment := &Environment{
		Packages:        map[string]*Package{},
		Functions:       map[string]*Function{},
		Types:           types.Default,
		StandardLibrary: standardLibrary,
	}

	return environment, nil
}

// ImportDirectory imports a directory to the environment.
func (env *Environment) ImportDirectory(pkg *Package) error {
	functions, structs, imports, errors := FindFunctions(pkg, env)
	return env.Import(pkg, functions, structs, imports, errors)
}

// Import imports the given functions and imports to the environment.
func (env *Environment) Import(pkg *Package, functions <-chan *Function, structs <-chan *types.Type, imports <-chan *Import, errors <-chan error) error {
	for {
		select {
		case err, ok := <-errors:
			if ok {
				return err
			}

		case imp, ok := <-imports:
			if !ok {
				continue
			}

			if env.Packages[imp.Path] != nil {
				continue
			}

			lastDotPosition := strings.LastIndex(imp.Path, ".")

			importPackage := &Package{
				Name: imp.Path[lastDotPosition+1:],
				Path: imp.FullPath,
			}

			env.Packages[imp.Path] = importPackage
			err := env.ImportDirectory(importPackage)

			if err != nil {
				return err
			}

		case typ, ok := <-structs:
			if !ok {
				return nil
			}

			if pkg.Name != MainPackageName {
				typ.Name = pkg.Name + "." + typ.Name
			}

			env.Types[typ.Name] = typ

		case function, ok := <-functions:
			if !ok {
				return nil
			}

			if pkg.Name != MainPackageName {
				function.Name = pkg.Name + "." + function.Name
			}

			env.Functions[function.Name] = function
		}
	}
}

// Compile compiles all functions.
func (env *Environment) Compile(optimize bool, verbose bool) {
	wg := sync.WaitGroup{}

	for _, function := range env.Functions {
		wg.Add(1)

		go func(function *Function) {
			defer wg.Done()
			Compile(function, env, optimize, verbose)

			if function.Error != nil {
				return
			}

			if atomic.AddInt64(&function.File.functionCount, -1) == 0 {
				function.File.Close()
			}
		}(function)
	}

	wg.Wait()
}
