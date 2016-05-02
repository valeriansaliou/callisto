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
  "fmt"
  "strings"

  "github.com/go-gl/gl/v4.1-core/gl"
)

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
  var (
    status    int32
    logLength int32
  )

  vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
  if err != nil {
    return 0, err
  }

  fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
  if err != nil {
    return 0, err
  }

  program := gl.CreateProgram()

  gl.AttachShader(program, vertexShader)
  gl.AttachShader(program, fragmentShader)
  gl.LinkProgram(program)

  gl.GetProgramiv(program, gl.LINK_STATUS, &status)

  if status == gl.FALSE {
    gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

    log := strings.Repeat("\x00", int(logLength + 1))
    gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

    return 0, fmt.Errorf("failed to link program: %v", log)
  }

  gl.DeleteShader(vertexShader)
  gl.DeleteShader(fragmentShader)

  return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
  var (
    status    int32
    logLength int32
  )

  shader := gl.CreateShader(shaderType)
  csources, free := gl.Strs(source)

  gl.ShaderSource(shader, 1, csources, nil)
  free()
  gl.CompileShader(shader)

  gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

  if status == gl.FALSE {
    gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

    log := strings.Repeat("\x00", int(logLength + 1))
    gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

    return 0, fmt.Errorf("failed to compile %v: %v", source, log)
  }

  return shader, nil
}
