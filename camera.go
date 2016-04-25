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
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type CameraData struct {
  Camera        mgl32.Mat4
  CameraUniform int32
}

var CAMERA CameraData

func (camera_data *CameraData) moveEyeX(increment float32) {
  camera_data.Camera[0] += increment
}

func (camera_data *CameraData) moveEyeY(increment float32) {
  camera_data.Camera[1] += increment
}

func (camera_data *CameraData) moveEyeZ(increment float32) {
  camera_data.Camera[2] += increment
}

func (camera_data *CameraData) moveCenterX(increment float32) {
  camera_data.Camera[3] += increment
}

func (camera_data *CameraData) moveCenterY(increment float32) {
  camera_data.Camera[4] += increment
}

func (camera_data *CameraData) moveCenterZ(increment float32) {
  camera_data.Camera[5] += increment
}

func (camera_data *CameraData) moveUpX(increment float32) {
  camera_data.Camera[6] += increment
}

func (camera_data *CameraData) moveUpY(increment float32) {
  camera_data.Camera[7] += increment
}

func (camera_data *CameraData) moveUpZ(increment float32) {
  camera_data.Camera[8] += increment
}

func getCamera() (*CameraData) {
  return &CAMERA
}

func createCamera(program uint32) {
  CAMERA.Camera = mgl32.LookAtV(CAMERA_DEFAULT_EYE, CAMERA_DEFAULT_CENTER, CAMERA_DEFAULT_UP)
  CAMERA.CameraUniform = gl.GetUniformLocation(program, gl.Str("cameraUniform\x00"))
}

func updateCamera() {
  key_state := getEventKeyState()

  if key_state.Up == true {
    getCamera().moveEyeX(CAMERA_MOVE_CELERITY_FORWARD)
  }

  if key_state.Down == true {
    getCamera().moveEyeX(CAMERA_MOVE_CELERITY_BACKWARD)
  }

  if key_state.Left == true {
    getCamera().moveEyeY(CAMERA_MOVE_CELERITY_FORWARD)
  }

  if key_state.Right == true {
    getCamera().moveEyeY(CAMERA_MOVE_CELERITY_BACKWARD)
  }
}

func bindCamera() {
  gl.UniformMatrix4fv(CAMERA.CameraUniform, 1, false, &CAMERA.Camera[0])
}
