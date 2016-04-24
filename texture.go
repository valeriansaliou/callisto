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
  "image"
  "image/draw"
  _ "image/jpeg"
  "os"

  "github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
  Id   int32
  Data uint32
}

var TEXTURE_ID_MAX int32 = 0

func loadTexture(name string) (Texture, error) {
  var texture = Texture{}

  filePath := fmt.Sprintf("assets/%s.jpg", name)

  imgFile, err := os.Open(filePath)
  if err != nil {
    return texture, fmt.Errorf("texture %q not found on disk: %v", filePath, err)
  }
  img, _, err := image.Decode(imgFile)
  if err != nil {
    return texture, err
  }

  rgba := image.NewRGBA(img.Bounds())
  if rgba.Stride != rgba.Rect.Size().X*4 {
    return texture, fmt.Errorf("unsupported stride")
  }
  draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

  // Assign an unique texture identifier
  TEXTURE_ID_MAX += 1
  texture.Id = TEXTURE_ID_MAX

  gl.ActiveTexture(uint32(texture.Id))
  gl.BindTexture(gl.TEXTURE_2D, texture.Data);

  gl.TexImage2D(
    gl.TEXTURE_2D,
    0,
    gl.RGBA,
    int32(rgba.Rect.Size().X),
    int32(rgba.Rect.Size().Y),
    0,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    gl.Ptr(rgba.Pix))

  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
  gl.GenerateMipmap(gl.TEXTURE_2D);
  gl.BindTexture(gl.TEXTURE_2D, 0);

  return texture, nil
}
