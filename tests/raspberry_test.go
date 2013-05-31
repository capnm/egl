// +build raspberry

package test

var (
	configAttr = []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.ALPHA_SIZE, 8,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
	contextAttr = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
	dstRect, srcRect            egl.VCRect
	NativeWindow                egl.EGLDispmanxWindow
	screenWidth, screenHeight uint32
)


func initPlatform() {
	screenWidth, screenHeight = egl.GraphicsGetDisplaySize(0)

	dstRect.X = 0
	dstRect.Y = 0
	dstRect.Width = int32(screenWidth)
	dstRect.Height = int32(screenHeight)

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.Width = int32(screenWidth << 16)
	srcRect.Height = int32(screenHeight << 16)

	dispman_display := egl.VCDispmanxDisplayOpen(0 /* LCD */)
	dispman_update := egl.VCDispmanxUpdateStart(0)

	dispman_element := egl.VCDispmanxElementAdd(
		dispman_update,
		dispman_display,
		0, /*layer */
		&dstRect,
		0, /*src */
		&srcRect,
		egl.DISPMANX_PROTECTION_NONE,
		nil, /*alpha */
		nil, /*clamp */
		0 /*transform */)

	NativeWindow.Element = dispman_element
	NativeWindow.Width = int(screen_width)
	NativeWindow.Height = int(screen_height)
	egl.VCDispmanxUpdateSubmitSync(dispman_update)
}

