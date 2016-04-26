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

type MatrixStack struct {
  Current *mgl32.Mat4
  Stack   *list.List
}

type MatrixUniforms struct {
  Model   int32
  Normal  int32
  Texture int32
}

var __MATRIX_STACK MatrixStack
var __MATRIX_UNIFORMS MatrixUniforms

func initializeMatrix() {
  identity_matrix := mgl32.Ident4()

  setMatrixStack(list.New())
  setMatrix(&identity_matrix)
}

func getMatrix() (*mgl32.Mat4) {
  return __MATRIX_STACK.Current
}

func setMatrix(matrix *mgl32.Mat4) {
  __MATRIX_STACK.Current = matrix
}

func getMatrixStack() (*list.List) {
  return __MATRIX_STACK.Stack
}

func setMatrixStack(matrix_stack *list.List) {
  __MATRIX_STACK.Stack = matrix_stack
}

func pushMatrix() {
  // Stack current matrix
  current_matrix := *getMatrix()

  getMatrixStack().PushBack(current_matrix)

  // Generate new current matrix
  new_matrix := mgl32.Ident4().Mul4(current_matrix)

  setMatrix(&new_matrix)
}

func popMatrix() {
  if getMatrixStack().Len() == 0 {
    panic("Cannot pop: matrix stack is empty")
  }

  last_element_list := getMatrixStack().Back()

  if previous_matrix, ok := (last_element_list.Value).(mgl32.Mat4); ok {
    // Assign now-popped element
    setMatrix(&previous_matrix)

    // Remove this element from the stack
    getMatrixStack().Remove(last_element_list)
  } else {
    panic("Cannot pop: error")
  }
}

func getMatrixUniforms() (*MatrixUniforms) {
  return &__MATRIX_UNIFORMS
}

func setMatrixUniforms(program uint32) {
  matrix_uniforms := getMatrixUniforms()

  matrix_uniforms.Model = gl.GetUniformLocation(program, gl.Str("modelUniform\x00"))
  matrix_uniforms.Normal = gl.GetUniformLocation(program, gl.Str("normalUniform\x00"))
  matrix_uniforms.Texture = gl.GetUniformLocation(program, gl.Str("textureUniform\x00"))
}
