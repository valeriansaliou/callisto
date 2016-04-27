/* Callisto - Yet another Solar System simulator
 *
 * Copyright (c) 2016, Valerian Saliou <valerian@valeriansaliou.name>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
  "container/list"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

// MatrixStack  Maps the stack of rendering matrixes
type MatrixStack struct {
  Current *mgl32.Mat4
  Stack   *list.List
}

// MatrixUniforms  Maps uniform matrixes
type MatrixUniforms struct {
  Model   int32
  Normal  int32
  Texture int32
}

// InstanceMatrixStack  Stores the stack of rendering matrixes
var InstanceMatrixStack MatrixStack

// InstanceMatrixUniforms  Stores uniform matrixes
var InstanceMatrixUniforms MatrixUniforms

func initializeMatrix() {
  identityMatrix := mgl32.Ident4()

  setMatrixStack(list.New())
  setMatrix(&identityMatrix)
}

func getMatrix() (*mgl32.Mat4) {
  return InstanceMatrixStack.Current
}

func setMatrix(matrix *mgl32.Mat4) {
  InstanceMatrixStack.Current = matrix
}

func getMatrixStack() (*list.List) {
  return InstanceMatrixStack.Stack
}

func setMatrixStack(matrixStack *list.List) {
  InstanceMatrixStack.Stack = matrixStack
}

func pushMatrix() {
  // Stack current matrix
  currentMatrix := *getMatrix()

  getMatrixStack().PushBack(currentMatrix)

  // Generate new current matrix
  newMatrix := mgl32.Ident4().Mul4(currentMatrix)

  setMatrix(&newMatrix)
}

func popMatrix() {
  if getMatrixStack().Len() == 0 {
    panic("Cannot pop: matrix stack is empty")
  }

  lastElementList := getMatrixStack().Back()

  if previousMatrix, ok := (lastElementList.Value).(mgl32.Mat4); ok {
    // Assign now-popped element
    setMatrix(&previousMatrix)

    // Remove this element from the stack
    getMatrixStack().Remove(lastElementList)
  } else {
    panic("Cannot pop: error")
  }
}

func getMatrixUniforms() (*MatrixUniforms) {
  return &InstanceMatrixUniforms
}

func setMatrixUniforms(program uint32) {
  matrixUniforms := getMatrixUniforms()

  matrixUniforms.Model = gl.GetUniformLocation(program, gl.Str("modelUniform\x00"))
  matrixUniforms.Normal = gl.GetUniformLocation(program, gl.Str("normalUniform\x00"))
  matrixUniforms.Texture = gl.GetUniformLocation(program, gl.Str("textureUniform\x00"))
}
