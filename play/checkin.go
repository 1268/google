package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (d Device) android_ID() (uint64, error) {
   v, ok := d.m.Fixed64(7) // long androidId_
   if ok {
      return v, nil
   }
   return 0, errors.New("androidId_")
}

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

// These can use default values, but they must all be included
type Config struct {
   Touch_Screen uint64 // 1
   Keyboard uint64 // 2
   Navigation uint64 // 3
   Screen_Layout uint64 // 4
   Has_Hard_Keyboard bool // 5
   Has_Five_Way_Navigation bool // 6
   Screen_Density uint64 // 7
   GL_ES_Version uint64 // 8
   System_Shared_Library []string // 9
   Platform string // 11
   GL_Extension []string // 15
   System_Available_Feature []string // 26
}

var Phone = Config{
   System_Available_Feature: []string{
      // app.source.getcontact
      "android.hardware.location.gps",
      // br.com.rodrigokolb.realdrum
      "android.software.midi",
      // com.app.xt
      "android.hardware.camera.front",
      // com.cabify.rider
      "android.hardware.camera.flash",
      // com.clearchannel.iheartradio.controller
      "android.hardware.microphone",
      // com.google.android.apps.walletnfcrel
      "android.software.device_admin",
      // com.google.android.youtube
      "android.hardware.touchscreen",
      "android.hardware.wifi",
      // com.madhead.tos.zh
      "android.hardware.sensor.accelerometer",
      // com.miHoYo.GenshinImpact
      "android.hardware.opengles.aep",
      // com.pinterest
      "android.hardware.camera",
      "android.hardware.location",
      "android.hardware.screen.portrait",
      // com.supercell.brawlstars
      "android.hardware.touchscreen.multitouch",
      // com.sygic.aura
      "android.hardware.location.network",
      // com.xiaomi.smarthome
      "android.hardware.bluetooth",
      "android.hardware.bluetooth_le",
      "android.hardware.camera.autofocus",
      "android.hardware.usb.host",
      // kr.sira.metal
      "android.hardware.sensor.compass",
      // org.thoughtcrime.securesms
      "android.hardware.telephony",
      // org.videolan.vlc
      "android.hardware.screen.landscape",
   },
   System_Shared_Library: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
      // com.miui.weather2
      "global-miui11-empty.jar",
   },
   GL_Extension: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   // com.axis.drawingdesk.v3
   // valid range 0x3_0001 - 0x7FFF_FFFF
   GL_ES_Version: 0xF_FFFF,
}

func (c Config) Checkin() ([]byte, error) {
   var m protobuf.Message
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 10, Type: 0, Value: protobuf.Varint(28)},
      }},
   }},
   protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(3)},
   protobuf.Field{Number: 18, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(490)},
      protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(196610)},
      // com.amctve.amcfullepisodes
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("org.apache.http.legacy")},
      // com.binance.dev
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.runner")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("x86")},
      // com.instagram.android
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
      // com.kakaogames.twodin
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
      // app.source.getcontact
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.gps")},
      }},
      // br.com.rodrigokolb.realdrum
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.midi")},
      }},
      // com.app.xt
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.front")},
      }},
      // com.cabify.rider
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.flash")},
      }},
      // com.clearchannel.iheartradio.controller
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
      }},
      // com.google.android.apps.walletnfcrel
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
      }},
      // com.google.android.youtube
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi")},
      }},
      // com.madhead.tos.zh
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
      }},
      // com.miHoYo.GenshinImpact
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.opengles.aep")},
      }},
      // com.pinterest
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
      }},
      // com.supercell.brawlstars
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
      }},
      // com.sygic.aura
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
      }},
      // com.xiaomi.smarthome
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth_le")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.host")},
      }},
      // kr.sira.metal
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
      }},
      // org.thoughtcrime.securesms
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony")},
      }},
      // org.videolan.vlc
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
      }},
   }},
   // r.Header.Set("User-Agent", "GoogleAuth/1.4 sargo PQ3B.190705.003")
   res, err := http.Post(
      "https://android.googleapis.com/checkin", "application/x-protobuffer",
      bytes.NewReader(checkin_body.Append(nil)),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}
