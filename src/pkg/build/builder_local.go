/*
  Copyright (c) 2018, Sylabs, Inc. All rights reserved.

  This software is licensed under a 3-clause BSD license.  Please
  consult LICENSE file distributed with the sources of this project regarding
  your rights to use or distribute this software.
*/
package build

import (
	"github.com/singularityware/singularity/src/pkg/image"
)

type LocalBuilder struct {
	Sandbox image.Sandbox
	Image   image.Image
	Definition
}

func NewLocalBuilder(j []byte) LocalBuilder {
	return LocalBuilder{image.Sandbox{}, &image.SIF{}, DefinitionFromJSON(j)}
}

func (*LocalBuilder) Build() {

}
