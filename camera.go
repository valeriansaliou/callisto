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
  "math"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

// CameraData  Maps camera state
type CameraData struct {
  Camera         mgl32.Mat4
  CameraUniform  int32

  PositionEye    mgl32.Vec3
  PositionTarget mgl32.Vec3

  InertiaDrag    float64
  InertiaTurn    float64

  ObjectIndex    int
  ObjectMatrix   mgl32.Mat4
  ObjectRadius   float32
  ObjectList     *[]string
}

// InstanceCamera  Stores camera state
var InstanceCamera CameraData

func (cameraData *CameraData) getOrbitObjectName() (string) {
  if cameraData.ObjectIndex > 0 {
    return (*cameraData.ObjectList)[cameraData.ObjectIndex - 1]
  }

  return ""
}

func (cameraData *CameraData) getEyeX() (position float32) {
  return cameraData.PositionEye[0]
}

func (cameraData *CameraData) getEyeY() (position float32) {
  return cameraData.PositionEye[1]
}

func (cameraData *CameraData) getEyeZ() (position float32) {
  return cameraData.PositionEye[2]
}

func (cameraData *CameraData) getTargetX() (position float32) {
  return cameraData.PositionTarget[0]
}

func (cameraData *CameraData) getTargetY() (position float32) {
  return cameraData.PositionTarget[1]
}

func (cameraData *CameraData) getTargetZ() (position float32) {
  return cameraData.PositionTarget[2]
}

func (cameraData *CameraData) moveEyeX(increment float32) {
  cameraData.PositionEye[0] += increment
}

func (cameraData *CameraData) moveEyeY(increment float32) {
  cameraData.PositionEye[1] += increment
}

func (cameraData *CameraData) moveEyeZ(increment float32) {
  cameraData.PositionEye[2] += increment
}

func (cameraData *CameraData) moveTargetX(increment float32) {
  cameraData.PositionTarget[0] += increment
}

func (cameraData *CameraData) moveTargetY(increment float32) {
  cameraData.PositionTarget[1] += increment
}

func (cameraData *CameraData) moveTargetZ(increment float32) {
  cameraData.PositionTarget[2] += increment
}

func (cameraData *CameraData) setObjectIndex(objectIndex int) {
  cameraData.ObjectIndex = objectIndex
}

func (cameraData *CameraData) defaultEye() {
  cameraData.PositionEye = ConfigCameraDefaultEye
}

func (cameraData *CameraData) defaultTarget() {
  cameraData.PositionTarget = ConfigCameraDefaultTarget
}

func (cameraData *CameraData) defaultInertia() {
  cameraData.InertiaDrag = 0.0
  cameraData.InertiaTurn = 0.0
}

func getCamera() (*CameraData) {
  return &InstanceCamera
}

func createCamera(program uint32) {
  camera := getCamera()

  camera.CameraUniform = gl.GetUniformLocation(program, gl.Str("cameraUniform\x00"))

  // Default inertia (none)
  camera.defaultInertia()

  // Default camera position
  camera.defaultEye()
  camera.defaultTarget()
}

func produceInertia(inertia *float64, increment float64, celerity float64) {
  *inertia += increment * celerity

  // Cap inertia to maximum value
  if *inertia > celerity {
    *inertia = celerity
  } else if *inertia < -1.0 * celerity {
    *inertia = -1.0 * celerity
  }
}

func consumeInertia(inertia *float64) (float64) {
  if *inertia > 0 {
    *inertia += ConfigCameraInertiaConsumeForward
  } else if *inertia < 0 {
    *inertia += ConfigCameraInertiaConsumeBackward
  }

  return *inertia
}

func processEventCameraEye() {
  camera := getCamera()

  if camera.ObjectIndex == 0 {
    var (
      celerity    float64
      rotationX   float64
      rotationY   float64

      inertiaDrag float64
      inertiaTurn float64
    )

    // Free flight camera
    keyState := getEventKeyState()
    timeFactor := normalizedTimeFactor()

    // Decrease speed if diagonal move
    if keyState.MoveTurbo == true {
      celerity = ConfigCameraMoveCelerityTurbo
    } else {
      celerity = ConfigCameraMoveCelerityCruise
    }

    if (keyState.MoveUp == true || keyState.MoveDown == true) && (keyState.MoveLeft == true || keyState.MoveRight == true) {
      celerity /= math.Sqrt(2.0)
    }

    // Acquire rotation around axis
    rotationX = float64(camera.getTargetX())
    rotationY = float64(camera.getTargetY())

    // Process camera move position (keyboard)
    if keyState.MoveUp == true {
      produceInertia(&(camera.InertiaDrag), ConfigCameraInertiaProduceForward, celerity)
    }
    if keyState.MoveDown == true {
      produceInertia(&(camera.InertiaDrag), ConfigCameraInertiaProduceBackward, celerity)
    }
    if keyState.MoveLeft == true {
      produceInertia(&(camera.InertiaTurn), ConfigCameraInertiaProduceForward, celerity)
    }
    if keyState.MoveRight == true {
      produceInertia(&(camera.InertiaTurn), ConfigCameraInertiaProduceBackward, celerity)
    }

    // Apply new position with inertia
    inertiaDrag = consumeInertia(&(camera.InertiaDrag))
    inertiaTurn = consumeInertia(&(camera.InertiaTurn))

    camera.moveEyeX(timeFactor * float32(inertiaDrag * -1.0 * math.Sin(rotationY) + inertiaTurn * math.Cos(rotationY)))
    camera.moveEyeZ(timeFactor * float32(inertiaDrag * math.Cos(rotationY) + inertiaTurn * math.Sin(rotationY)))
    camera.moveEyeY(timeFactor * float32(inertiaDrag * math.Sin(rotationX)))
  } else {
    // Orbit camera
    size := normalizeObjectSize(camera.ObjectRadius)

    camera.PositionEye = mgl32.Vec3{0, 0, -1 * size * ConfigCameraOrbitMagnification}
  }

  // Translation: walk
  camera.Camera = camera.Camera.Mul4(mgl32.Translate3D(camera.getEyeX(), camera.getEyeY(), camera.getEyeZ()))
}

func processEventCameraTarget() {
  camera := getCamera()
  keyState := getEventKeyState()
  timeFactor := normalizedTimeFactor()

  camera.moveTargetX(timeFactor * keyState.WatchY * float32(math.Pi) * ConfigCameraTargetAmortizeFactor)
  camera.moveTargetY(timeFactor * keyState.WatchX * float32(math.Pi) * 2 * ConfigCameraTargetAmortizeFactor)

  // Rotation: view
  camera.Camera = camera.Camera.Mul4(mgl32.HomogRotate3D(camera.getTargetX(), mgl32.Vec3{1, 0, 0}))
  camera.Camera = camera.Camera.Mul4(mgl32.HomogRotate3D(camera.getTargetY(), mgl32.Vec3{0, 1, 0}))
  camera.Camera = camera.Camera.Mul4(mgl32.HomogRotate3D(camera.getTargetZ(), mgl32.Vec3{0, 0, 1}))
}

func updateCamera() {
  camera := getCamera()

  // Update overall camera position
  if camera.ObjectIndex == 0 {
    // Free flight
    camera.Camera = mgl32.Ident4()
  } else {
    // Orbit flight
    camera.Camera = camera.ObjectMatrix
  }

  // Orbit camera or free flight camera? (reverse rotation <> translation)
  if camera.ObjectIndex == 0 {
    // Free flight camera
    processEventCameraTarget()
    processEventCameraEye()
  } else {
    // Orbit camera
    processEventCameraEye()
    processEventCameraTarget()
  }
}

func toggleNextCameraObject() {
  camera := getCamera()

  // Go to next index
  camera.ObjectIndex++

  // Index overflow?
  if camera.ObjectIndex > len(*camera.ObjectList) {
    camera.ObjectIndex = 0
  }

  // Reset camera state
  resetCamera()
}

func resetCamera() {
  camera := getCamera()

  // Reset camera modifiers
  resetMouseCursor()

  // Reset camera itself
  camera.defaultInertia()

  camera.defaultEye()
  camera.defaultTarget()
}

func resetCameraObject() {
  getCamera().ObjectIndex = 0
}

func initializeCameraLocks(objects *[]Object) {
  // Initialize object list storage space
  var objectList []string

  getCamera().ObjectList = &objectList

  // Create camera locks (in object list)
  createCameraLocks(objects)
}

func createCameraLocks(objects *[]Object) {
  camera := getCamera()

  for o := range *objects {
    *camera.ObjectList = append(*camera.ObjectList, (*objects)[o].Name)

    // Create locks for child objects
    createCameraLocks(&((*objects)[o]).Objects)
  }
}

func bindCamera() {
  camera := getCamera()

  gl.UniformMatrix4fv(camera.CameraUniform, 1, false, &(camera.Camera[0]))
}
