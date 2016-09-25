package backend

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/acodec"
	"github.com/dradtke/go-allegro/allegro/audio"
	"github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/go-allegro/allegro/font/ttf"
	"github.com/dradtke/go-allegro/allegro/image"
	"github.com/dradtke/go-allegro/allegro/primitives"
)

var (
	isKeyboardInited    = false
	isMouseInited       = false
	isJoystickInited    = false
	isAudioInited       = false
	isAudioCodecsInited = false
	isFontInited        = false
	isTTFInited         = false
	isImageIOInited     = false
	isPrimitivesInited  = false
)

// InitKeyboard initializes Allegro's keyboard subsystem.
func InitKeyboard() error {
	if isKeyboardInited {
		return nil
	}

	if err := allegro.InstallKeyboard(); err != nil {
		return err
	}

	isKeyboardInited = true
	return nil
}

// InitMouse initializes Allegro's mouse subsystem.
func InitMouse() error {
	if isMouseInited {
		return nil
	}

	if err := allegro.InstallMouse(); err != nil {
		return err
	}

	isMouseInited = true
	return nil
}

// InitJoystick initializes Allegro's joystick subsystem.
func InitJoystick() error {
	if isJoystickInited {
		return nil
	}

	if err := allegro.InstallJoystick(); err != nil {
		return err
	}

	isJoystickInited = true
	return nil
}

// InitAudio initializes Allegro's audio subsystem.
func InitAudio() error {
	if isAudioInited {
		return nil
	}

	// TODO: I used to do this on Linux:
	//       al_set_config_value(al_get_system_config(), "audio", "driver", "alsa");
	//       'cause PulseAudio was used by default and didn't work well.

	if err := audio.Install(); err != nil {
		return err
	}

	isAudioInited = true
	return nil
}

// InitAudioCodecs initializes Allegro's audio codecs subsystem.
func InitAudioCodecs() error {
	if isAudioCodecsInited {
		return nil
	}

	if err := InitAudio(); err != nil {
		return err
	}

	if err := acodec.Install(); err != nil {
		return err
	}

	isAudioCodecsInited = true
	return nil
}

// InitFont initializes Allegro's font subsystem.
func InitFont() error {
	if isFontInited {
		return nil
	}

	font.Install() // Does not return anything
	isFontInited = true
	return nil
}

// InitTTF intializes Allegro's TTF subsystem.
func InitTTF() error {
	if isTTFInited {
		return nil
	}

	if err := InitFont(); err != nil {
		return err
	}

	ttf.Install() // TODO: Should return a success flag, at least in Allegro 5.2
	isTTFInited = true
	return nil
}

// InitImageIO initializes Allegro's image I/O subsystem.
func InitImageIO() error {
	if isImageIOInited {
		return nil
	}

	if err := image.Install(); err != nil {
		return err
	}

	isImageIOInited = true
	return nil
}

// InitPrimitives initializes Allegro's primitives subsystem.
func InitPrimitives() error {
	if isPrimitivesInited {
		return nil
	}

	if err := primitives.Install(); err != nil {
		return err
	}

	isPrimitivesInited = true
	return nil
}

// InitializeAllegro installs all Allegro modules.
func InitializeAllegro() {
	InitKeyboard()
	InitMouse()
	InitJoystick()
	InitAudio()
	InitAudioCodecs()
	InitFont()
	InitTTF()
	InitImageIO()
	InitPrimitives()
}

// ShutdownAllegro uninstalls all installed Allegro modules.
func ShutdownAllegro() {

	if isImageIOInited {
		image.Uninstall()
	}

	if isTTFInited {
		ttf.Uninstall()
	}

	if isJoystickInited {
		allegro.UninstallJoystick()
	}

	if isKeyboardInited {
		allegro.UninstallKeyboard()
	}

	if isMouseInited {
		allegro.UninstallMouse()
	}

	if isAudioInited {
		audio.Uninstall()
	}

	if isFontInited {
		font.Uninstall()
	}

	if isPrimitivesInited {
		primitives.Uninstall()
	}
}
