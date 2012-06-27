package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func DanishStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x44,0x97,
0x4d,0x92,0xdb,0x38,0x12,0x85,0xf7,0x79,0x8c,0x59,0xf5,0xa2,
0x42,0x77,0xe8,0xee,0xf2,0x84,0x67,0x1c,0x76,0x74,0x44,0xb9,
0x7b,0x0f,0x99,0x14,0x45,0x11,0x04,0x14,0x00,0xc8,0x19,0x1d,
0xa8,0x2a,0x62,0x6e,0x50,0x7b,0x5d,0x6c,0xbe,0x97,0x90,0xbb,
0x16,0x25,0x4b,0x24,0x90,0xbf,0xef,0xbd,0x4c,0xcf,0x96,0x27,
0x0b,0x27,0x1b,0x93,0x8d,0xc5,0xda,0x1c,0x2d,0x34,0xab,0x79,
0xb5,0x81,0x27,0xd7,0xfb,0xab,0x9d,0x72,0xb1,0x75,0x1c,0x6c,
0x0f,0xc5,0x8e,0x71,0xdc,0x79,0x51,0xf8,0x6b,0xfc,0x19,0x9f,
0xa7,0x12,0xec,0x1c,0x12,0x7f,0xc5,0xfe,0x65,0x3b,0x07,0xe7,
0x65,0x19,0x0d,0x03,0x2b,0x06,0xea,0x3c,0x61,0xbf,0x62,0xe6,
0xb3,0x9f,0xd9,0xb9,0x74,0xde,0xb1,0xf8,0xcc,0xcb,0x67,0xae,
0x8f,0x31,0x62,0x6e,0xe1,0xe5,0x78,0x6a,0x7c,0xab,0x73,0xb2,
0x2d,0xc9,0xc5,0xca,0xb3,0xbc,0xf3,0xe5,0x74,0x7f,0x2f,0xb5,
0x71,0x2f,0xa4,0x6a,0x43,0xb0,0x96,0xb9,0x69,0x6b,0x1e,0x4c,
0x76,0xef,0xaf,0x9c,0x98,0x17,0x45,0x35,0x56,0xdb,0xef,0x6f,
0x65,0xe4,0x7b,0xc4,0x3b,0x96,0x57,0x8b,0xf3,0x34,0x61,0x23,
0xa4,0x81,0xe7,0xcb,0x96,0x12,0x37,0x43,0x9a,0x38,0x43,0xda,
0xe7,0x2d,0xe9,0x59,0xbf,0x45,0x2a,0x84,0x32,0xda,0xa7,0xc4,
0xdd,0xc9,0xf3,0xd3,0xe9,0x6d,0xc0,0x67,0xaa,0x8b,0x05,0xec,
0x79,0x01,0x78,0x55,0xc7,0xa4,0xa3,0x0b,0x59,0x1e,0xe3,0xac,
0x20,0x57,0xfd,0x1e,0x13,0x31,0x2d,0x9b,0x0e,0xd6,0xb9,0x29,
0x5d,0xcc,0x79,0x5a,0xcf,0x4a,0xad,0x65,0xce,0xe4,0xab,0xd5,
0xb0,0x7a,0x6d,0xc2,0xda,0xa8,0xd3,0x52,0xe6,0x34,0xf5,0x67,
0x94,0x7f,0x20,0x60,0x99,0x72,0x9f,0x32,0x43,0x08,0xe7,0xb0,
0xf2,0x45,0xfd,0xa8,0x4b,0x88,0xf6,0x0b,0x05,0x19,0x1a,0x8e,
0x0f,0xe1,0xa0,0x46,0x50,0x9a,0x9d,0xb6,0x2d,0x78,0x6f,0x2a,
0x7a,0xe5,0x15,0xf9,0x36,0x6b,0x72,0x77,0xe2,0x51,0x25,0x62,
0x9c,0xdb,0x1f,0xd4,0x8b,0x37,0x4a,0xf5,0xf0,0xa5,0x1c,0xa8,
0xec,0xab,0x3a,0x7e,0x0c,0xd4,0x83,0x36,0xe3,0x30,0xca,0x3d,
0xce,0x27,0x3b,0xde,0x5f,0xe9,0x15,0x5d,0xab,0x64,0x1b,0x77,
0x9b,0x54,0x8d,0xd5,0xaf,0xd3,0x48,0x1c,0x2e,0x32,0x33,0xf3,
0xb3,0x1a,0x9f,0x96,0xf2,0x44,0xda,0xe7,0x91,0x8f,0x74,0x93,
0x5f,0xc2,0x4a,0x9b,0x0d,0x73,0xad,0xa3,0x4d,0x34,0x68,0xd3,
0x45,0xee,0x71,0xe0,0x07,0x81,0x97,0xfb,0x9b,0x70,0xd2,0x6e,
0x4a,0xf3,0x39,0xf4,0x3a,0xda,0x44,0x6f,0xec,0x9c,0xab,0x3c,
0x54,0xfb,0x4c,0x6f,0x9e,0x43,0x5a,0x43,0x59,0x04,0x16,0x2c,
0x52,0x06,0x41,0x01,0x87,0x42,0x6a,0x6d,0x0f,0x68,0x74,0x88,
0x08,0x03,0x42,0xeb,0x25,0xa4,0x0d,0x34,0xea,0xa1,0xad,0x1c,
0x75,0x44,0xac,0x42,0x0a,0x0d,0xe7,0x17,0x79,0xad,0xe1,0x62,
0xdf,0xcf,0x5e,0x70,0x3b,0x8e,0xd3,0x0d,0x7b,0xa3,0xfd,0x05,
0x78,0x8f,0x37,0xa2,0x54,0x56,0xc4,0xea,0x45,0xbc,0x6c,0x69,
0xee,0x8e,0x30,0x43,0x24,0xad,0x5a,0x12,0xe6,0xa0,0x84,0xf0,
0x55,0xc7,0x6b,0x1b,0xd7,0x23,0xed,0xcd,0x4b,0xcb,0xfa,0x37,
0x5c,0x0b,0xa1,0x35,0x8c,0x4f,0x05,0x77,0x56,0xaf,0xca,0x19,
0x9f,0x0d,0xc4,0x25,0x10,0xed,0xa7,0x4f,0x85,0xcb,0x29,0xec,
0xc9,0xc2,0x36,0x6d,0x1e,0xfb,0x8f,0xfe,0xe6,0xb2,0xc5,0xd9,
0x3e,0xa9,0xb6,0x38,0xfc,0x27,0xec,0x3a,0x8d,0xc7,0xa2,0x7c,
0x02,0x18,0x9b,0x61,0x8a,0x0a,0xf6,0x55,0x60,0x68,0x61,0x80,
0x5e,0x32,0x42,0x2b,0xa8,0x7c,0xb3,0x17,0x01,0xf2,0x66,0xbf,
0x9f,0xcb,0x5c,0xdb,0x0c,0x81,0x06,0xa0,0x42,0x1f,0xb2,0xf0,
0x1c,0xbd,0xf6,0x53,0x58,0xf9,0x84,0x7e,0x62,0xc2,0xfd,0x55,
0x45,0xc5,0xfc,0x44,0x17,0x7b,0xa0,0xc5,0xbe,0x8d,0xff,0x51,
0xed,0xc8,0x9d,0x62,0x1e,0x55,0x8b,0xa8,0x16,0x5e,0x72,0xc1,
0x54,0xc0,0x0b,0xb4,0x53,0x42,0xe0,0x0c,0x64,0xa9,0x16,0xdb,
0x30,0xd1,0xbc,0x26,0x3e,0x10,0x23,0xa1,0x38,0x2a,0xd2,0xcd,
0xe6,0x0a,0xb3,0x0c,0x48,0x91,0xe6,0xdf,0xd1,0x93,0x6a,0x75,
0x30,0xd0,0x39,0x61,0x31,0x82,0xc2,0xb3,0x13,0x1d,0x0f,0x33,
0x01,0xee,0xce,0x35,0xef,0xd3,0x42,0x85,0x3d,0x2c,0x41,0x76,
0x5c,0xea,0x01,0xe7,0xef,0x09,0xc7,0x0e,0x03,0x08,0xab,0x7f,
0xfe,0xa1,0x4e,0x4e,0x0a,0xa4,0x2e,0xc5,0xf9,0xe9,0x28,0xfe,
0xd3,0x65,0xe4,0x37,0x35,0xf4,0xcb,0xfd,0xfd,0x38,0x26,0xf8,
0x98,0x1c,0x61,0xd4,0x15,0x98,0x88,0x12,0x08,0x50,0x1a,0x04,
0x53,0xf2,0x15,0x7c,0x24,0x74,0x9f,0x5d,0x73,0xb0,0x96,0xf8,
0x41,0x4c,0xe7,0x1c,0x07,0x81,0xe7,0x40,0x7e,0x5c,0x50,0xb5,
0xbc,0x34,0x2b,0xb2,0x15,0x3d,0xd1,0xfb,0x1b,0x61,0xc0,0x0a,
0xe5,0xbe,0x07,0x14,0x10,0x17,0xa8,0x08,0x79,0x1d,0x6f,0xc8,
0x0f,0xf1,0x08,0xd6,0x7b,0x4e,0xde,0x89,0x44,0xf7,0x06,0xfc,
0x7a,0xec,0x1f,0x91,0x55,0x9a,0x4d,0xca,0x14,0x49,0x88,0x24,
0xc6,0x5e,0xdb,0xe3,0xe8,0x12,0x46,0x2e,0x3f,0xa9,0x50,0xed,
0xdf,0xf9,0x4c,0xf5,0xa4,0x43,0x54,0xd0,0x6a,0xdc,0x5a,0xea,
0x2e,0x11,0x63,0x87,0x87,0xc4,0xd0,0x75,0xe7,0x8f,0xd1,0x45,
0x55,0x65,0x51,0xbe,0xca,0xd1,0x09,0x5d,0x45,0x64,0x3b,0x39,
0xc2,0xbf,0x82,0x93,0x44,0xe9,0x25,0x9b,0x23,0x10,0x9a,0x28,
0x52,0xe6,0x39,0x8a,0xd3,0xec,0xcf,0x97,0x5f,0xa1,0xb1,0x0b,
0xa4,0x42,0x98,0xe6,0x53,0xe3,0x43,0xdd,0xaa,0x37,0x0c,0xcc,
0x11,0x49,0x88,0x43,0x01,0x87,0xf1,0xfe,0x46,0x0c,0x60,0x36,
0xce,0x82,0x0e,0x84,0x21,0xbd,0x50,0x8e,0xe3,0xe5,0x23,0xf4,
0x83,0xa1,0xac,0xc9,0xcb,0x81,0xfd,0x15,0xb5,0x1b,0x8b,0x28,
0x18,0x55,0x05,0x49,0x5a,0x26,0x27,0x64,0x71,0x16,0xb0,0x9f,
0x5d,0x3a,0xbc,0xc6,0xd8,0x83,0x2e,0xae,0x70,0x82,0xeb,0xdc,
0x66,0x65,0x79,0x0e,0x04,0xb3,0x7a,0x86,0x6a,0x8a,0xdd,0xff,
0xa7,0x81,0x83,0xb7,0x6b,0x8e,0x8f,0x23,0x02,0x82,0x86,0x19,
0x94,0x15,0x9c,0x68,0xab,0xdc,0x38,0xe1,0x85,0x6a,0xce,0xab,
0x4b,0xde,0xe1,0x86,0x0a,0xa2,0x99,0xe8,0x11,0xbd,0x99,0x78,
0xcb,0x75,0x82,0xba,0x09,0x44,0x79,0x2d,0x92,0x43,0x52,0x27,
0xad,0xd1,0xbe,0x43,0xf0,0xbc,0x3e,0xf9,0x2f,0x55,0x4d,0xcc,
0x26,0xe1,0x48,0xd3,0x54,0x48,0x47,0xdf,0x4e,0x20,0x1c,0x45,
0x03,0xae,0x57,0xda,0x0c,0x33,0x36,0x18,0x47,0xe2,0x65,0x3c,
0x1a,0x48,0x1e,0xd7,0x2b,0x53,0x4a,0x51,0xd8,0x67,0xc8,0x9a,
0xcb,0x3c,0x1e,0x54,0x26,0xdc,0x8b,0x72,0x0a,0x49,0x83,0x51,
0x8c,0xea,0x42,0x12,0x03,0x35,0xcb,0x57,0x2f,0x2c,0xda,0x14,
0xf9,0x45,0x0b,0xdc,0x17,0x17,0x1c,0x68,0x12,0x3a,0x54,0x63,
0xf1,0x89,0xe0,0x59,0x69,0x54,0x69,0x54,0x8e,0x05,0xe7,0xc0,
0xe9,0x89,0x99,0x47,0xc5,0xb9,0x3b,0x04,0xa8,0x5f,0xec,0xc5,
0xe7,0xca,0x98,0x90,0x7e,0x6e,0x2d,0x73,0xcc,0xbd,0xa0,0xaa,
0x86,0x66,0xc3,0x93,0x24,0xb2,0x83,0x59,0x6d,0x7c,0x30,0xd7,
0x75,0xdf,0xce,0x97,0xfb,0x5b,0xbc,0xda,0x35,0x86,0xa1,0xd2,
0x27,0x12,0x92,0x26,0xb9,0x04,0xb4,0xfb,0x1b,0x41,0x6a,0xf0,
0x5c,0x41,0x34,0xfe,0xd9,0x25,0xfa,0x17,0x7b,0xe9,0xb3,0x53,
0x8c,0x77,0x98,0x50,0x35,0x19,0xdb,0xe7,0x60,0x25,0x4b,0xa1,
0x50,0x13,0x97,0x00,0x01,0x03,0x7b,0xc0,0xc8,0x81,0xc7,0x74,
0x08,0xa9,0x31,0xff,0x18,0x57,0xa2,0x5f,0xab,0xca,0x40,0x03,
0x0f,0x26,0x2d,0xe2,0x6c,0xbc,0xf1,0x80,0xb3,0x5d,0x31,0x94,
0xae,0x77,0xec,0x55,0x43,0xf2,0xc9,0xbe,0xd3,0x59,0xcd,0xc5,
0x87,0x72,0x2b,0xca,0x94,0x17,0x86,0x64,0x26,0xfa,0x95,0x1b,
0x73,0x26,0xb0,0xbe,0x86,0x7c,0x84,0x85,0xd6,0x09,0x81,0x38,
0xca,0xe8,0xba,0x50,0xa3,0x82,0xb3,0x17,0x79,0x61,0xea,0x75,
0x4b,0x0b,0x48,0xd6,0xbc,0xd0,0xe4,0xb9,0xbf,0x93,0xb1,0xd0,
0x4d,0x53,0x35,0xc0,0x42,0xda,0xbb,0x16,0x46,0x56,0x81,0x26,
0xa1,0xd0,0xae,0xb3,0x5e,0xbb,0x44,0x6a,0xce,0xfa,0xe0,0x65,
0x23,0x19,0x68,0x32,0xca,0x41,0x99,0x23,0xa5,0xa8,0x04,0xa2,
0xe0,0xff,0xfa,0xea,0x45,0x74,0x9f,0xbd,0x14,0x89,0xce,0x68,
0xbd,0xc0,0x12,0x4f,0xed,0xca,0x5c,0x60,0xcc,0x0f,0x17,0xa0,
0xfe,0x98,0xec,0x0a,0x06,0x42,0x04,0x6c,0x48,0x79,0x24,0x29,
0xae,0xc0,0x9a,0xde,0x1c,0x20,0x57,0xa7,0xbe,0x96,0x2c,0x5e,
0xdf,0x18,0xc2,0x0a,0xb0,0x02,0xc9,0x6e,0xc2,0x04,0xb6,0x66,
0xbf,0x87,0x12,0xed,0x9b,0x0f,0xf5,0x01,0xda,0x0a,0xbb,0x9d,
0xcb,0xa2,0xd1,0x7f,0x8d,0xce,0xbb,0x04,0x3f,0x98,0x41,0xbb,
0x8e,0xdb,0xca,0xe6,0x46,0xa0,0x9a,0x0a,0x28,0x4d,0x8f,0x87,
0x95,0xac,0xb3,0x96,0x35,0x47,0xd4,0x67,0xa0,0x25,0x7d,0x95,
0xea,0x28,0x29,0xde,0x8a,0x68,0xdf,0xcf,0x79,0x0d,0x4c,0x7d,
0xd1,0x4b,0xc1,0x7b,0x70,0x5f,0xe6,0xc2,0x2b,0x87,0xba,0x16,
0x85,0x2b,0x8c,0xb6,0x6f,0x2e,0xcc,0x2e,0xb8,0x3f,0x69,0x6a,
0x5f,0xe7,0x1f,0xe7,0x00,0x99,0x00,0x84,0x2b,0x12,0x6e,0x3e,
0x9a,0x05,0x0c,0x7b,0xaa,0xa7,0x40,0x6f,0x9c,0xd0,0x12,0x4b,
0x50,0xa6,0xed,0x0d,0xb1,0xd1,0x6c,0xc1,0xef,0xc5,0x56,0x06,
0x2f,0xc7,0x45,0xf8,0xc7,0x12,0xa6,0xf5,0x53,0xdc,0xd6,0xec,
0xd3,0x5d,0x6d,0x12,0x5a,0x10,0xa8,0x54,0x06,0x15,0xe0,0x80,
0x2d,0xc6,0x37,0x33,0x70,0x5e,0xe6,0x95,0xa6,0xe0,0x0e,0x36,
0x43,0x49,0xed,0xa7,0x5c,0xef,0xdb,0x52,0x2e,0x6c,0xd8,0xe1,
0xf4,0x44,0xc8,0x6a,0x21,0x22,0xb1,0x6a,0x90,0xf2,0xa7,0x15,
0x56,0xe0,0x00,0x43,0xbe,0xab,0x10,0xe3,0xc7,0x48,0x38,0xfc,
0xad,0x64,0x58,0x9d,0xa0,0xad,0x2c,0x7f,0xda,0x4a,0xbe,0xb2,
0x05,0xcf,0xf1,0xa4,0x26,0x48,0x21,0x0b,0xe2,0x21,0xad,0xf1,
0x21,0xcc,0xd8,0x78,0x2b,0x80,0x32,0xdd,0x34,0x07,0x77,0xef,
0x14,0x2f,0x7f,0xd5,0x8a,0x2f,0x96,0xbb,0xa6,0x48,0x8b,0x7c,
0xec,0x1b,0xdc,0x85,0xf5,0x8f,0xb5,0x0f,0x90,0xf8,0x72,0x4b,
0xff,0x3f,0xa5,0xc9,0xd9,0xe2,0xe8,0x47,0xb6,0xb5,0x87,0x92,
0xc7,0x63,0x4b,0x17,0x51,0x24,0x97,0x0f,0x5d,0x7f,0x7a,0xa4,
0x0a,0xe2,0x3b,0x48,0xb4,0x13,0xfc,0x34,0x07,0x3e,0x7c,0xf7,
0x5f,0x0f,0x9a,0x90,0x7d,0xb9,0xef,0x83,0xca,0x96,0xb8,0x1d,
0xc9,0xd5,0x57,0x54,0x1f,0x78,0xef,0xa2,0xb4,0x6b,0x5f,0xd7,
0x42,0x36,0xd0,0x31,0x69,0x24,0xd2,0x72,0xd6,0x1c,0xad,0x1c,
0x2a,0xd1,0x88,0x90,0x49,0xa1,0xc4,0x12,0x1f,0xbd,0x2f,0xf8,
0x57,0x0a,0xa7,0x1c,0x19,0x9f,0x4a,0x8c,0x51,0x2c,0xb6,0xfd,
0x00,0x43,0x41,0x70,0xd0,0xfa,0xd8,0x9b,0x03,0x4b,0x77,0x3f,
0xec,0xf8,0x3a,0xf0,0xb2,0xe7,0xf8,0x5b,0xdf,0x9a,0x1f,0x99,
0x38,0x48,0xcf,0x0f,0x65,0xee,0x7b,0x67,0x5f,0x58,0x8a,0x7d,
0xf2,0x19,0xad,0x35,0xa0,0x0f,0x98,0xbe,0x0c,0xfd,0x82,0xe2,
0x52,0x20,0xe5,0xea,0x48,0x3a,0x33,0x9a,0xbf,0xf5,0xdd,0x8c,
0xff,0x0b,0x31,0x13,0x09,0x93,0x69,0x7a,0x90,0x1c,0xfb,0xc0,
0x80,0xba,0x0f,0xf6,0x80,0x78,0x80,0x8f,0xac,0xe6,0x12,0x81,
0xcd,0x63,0x58,0x1c,0xc0,0x04,0xc5,0x4e,0xa1,0xa1,0x4c,0x42,
0x1b,0xaf,0x9f,0xec,0xff,0x01,0x00,0x00,0xff,0xff,0x25,0xfb,
0xc5,0xbf,0xb0,0x0d,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}