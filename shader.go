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

var vertexShader = `
#version 330

uniform mat4 projectionUniform;
uniform mat4 cameraUniform;
uniform mat4 modelUniform;

in vec3 vertAttrib;
in mat3 normalUniform;

out vec3 N;
out vec4 point;

void main() {
  point = modelUniform * vec4(vertAttrib, 1.0);
  N = normalUniform * vertAttrib;

  gl_Position = projectionUniform * cameraUniform * modelUniform * vec4(vertAttrib, 1);
}
` + "\x00"

var fragmentShader = `
#version 330

precision mediump float;

const vec3 Kd = vec3(0.4, 0.9, 0.1);
const vec3 Ks = vec3(0.9, 0.9, 0.9);
const float ns = 80.0;

uniform vec4 lighting;

in vec3 N;
in vec4 point;

out vec4 color;

void main() {
  vec3 N1 = normalize(N);

  vec3 L1 = normalize((lighting - point).xyz);

  float NL = clamp(dot(N1, L1), 0.0, 1.0);
  color = vec4(Kd * NL, 1.0);

  vec3 mV = normalize(point.xyz);
  vec3 R = reflect(mV, N1);
  float RL = clamp(dot(R, L1), 0.0, 1.0);
  color += vec4(Ks * pow(RL, ns), 0.0);
}
` + "\x00"
