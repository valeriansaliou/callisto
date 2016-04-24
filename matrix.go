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

var (
  CURRENT_MATRIX  mgl32.Mat4
  MATRIX_STACK    *list.List

  MODEL_UNIFORM   int32
  TEXTURE_UNIFORM int32
)

func initializeMatrix() {
  MATRIX_STACK = list.New()
  CURRENT_MATRIX = mgl32.Mat4{}
}

func getMatrix() (*mgl32.Mat4) {
  return &CURRENT_MATRIX
}

func pushMatrix() {
  CURRENT_MATRIX = mgl32.Ident4()
  MATRIX_STACK.PushFront(CURRENT_MATRIX)
}

func popMatrix() {
  if MATRIX_STACK.Len() == 0 {
    panic("Cannot pop: matrix stack is empty")
  }

  MATRIX_STACK.Remove(MATRIX_STACK.Front())
}

func setMatrixUniforms(program uint32) {
  MODEL_UNIFORM = gl.GetUniformLocation(program, gl.Str("modelUniform\x00"))
  TEXTURE_UNIFORM = gl.GetUniformLocation(program, gl.Str("textureUniform\x00"))
}
