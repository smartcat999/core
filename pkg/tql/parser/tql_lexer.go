// AUTOGENERATED FILE

// +build !codeanalysis
// Generated from /home/jesse/myspace/src/tkeel-io/core/pkg/tql/TQL.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 51, 532,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 193, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 204, 10, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 212, 10, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 5, 14, 218, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 5, 15, 226, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	5, 16, 235, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 242, 10,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43,
	3, 43, 6, 43, 381, 10, 43, 13, 43, 14, 43, 382, 3, 43, 7, 43, 386, 10,
	43, 12, 43, 14, 43, 389, 11, 43, 3, 43, 7, 43, 392, 10, 43, 12, 43, 14,
	43, 395, 11, 43, 3, 43, 6, 43, 398, 10, 43, 13, 43, 14, 43, 399, 5, 43,
	402, 10, 43, 3, 44, 3, 44, 3, 44, 7, 44, 407, 10, 44, 12, 44, 14, 44, 410,
	11, 44, 3, 45, 3, 45, 7, 45, 414, 10, 45, 12, 45, 14, 45, 417, 11, 45,
	3, 46, 3, 46, 3, 46, 7, 46, 422, 10, 46, 12, 46, 14, 46, 425, 11, 46, 5,
	46, 427, 10, 46, 3, 47, 5, 47, 430, 10, 47, 3, 47, 3, 47, 3, 48, 5, 48,
	435, 10, 48, 3, 48, 6, 48, 438, 10, 48, 13, 48, 14, 48, 439, 3, 48, 3,
	48, 6, 48, 444, 10, 48, 13, 48, 14, 48, 445, 3, 48, 6, 48, 449, 10, 48,
	13, 48, 14, 48, 450, 3, 48, 3, 48, 3, 48, 3, 48, 6, 48, 457, 10, 48, 13,
	48, 14, 48, 458, 5, 48, 461, 10, 48, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49,
	467, 10, 49, 12, 49, 14, 49, 470, 11, 49, 3, 49, 3, 49, 3, 50, 6, 50, 475,
	10, 50, 13, 50, 14, 50, 476, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52,
	3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3,
	58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63,
	3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3,
	68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73,
	3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 2, 2, 77, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24,
	47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33,
	65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42,
	83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51,
	101, 2, 103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2,
	119, 2, 121, 2, 123, 2, 125, 2, 127, 2, 129, 2, 131, 2, 133, 2, 135, 2,
	137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 147, 2, 149, 2, 151, 2, 3, 2, 39,
	8, 2, 37, 37, 44, 44, 50, 59, 67, 92, 97, 97, 99, 124, 7, 2, 37, 38, 47,
	47, 66, 92, 97, 97, 99, 124, 3, 2, 50, 59, 6, 2, 37, 37, 67, 92, 97, 97,
	99, 124, 8, 2, 37, 38, 47, 48, 50, 59, 66, 92, 97, 97, 99, 124, 7, 2, 37,
	37, 44, 44, 67, 92, 97, 97, 99, 124, 8, 2, 37, 38, 47, 47, 50, 59, 66,
	92, 97, 97, 99, 124, 3, 2, 51, 59, 4, 2, 45, 45, 47, 47, 3, 2, 41, 41,
	5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100,
	100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103,
	103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106,
	106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109,
	109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112,
	112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115,
	115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118,
	118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121,
	121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124,
	124, 2, 533, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2,
	2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2,
	2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3,
	2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 3, 153,
	3, 2, 2, 2, 5, 155, 3, 2, 2, 2, 7, 162, 3, 2, 2, 2, 9, 167, 3, 2, 2, 2,
	11, 170, 3, 2, 2, 2, 13, 174, 3, 2, 2, 2, 15, 179, 3, 2, 2, 2, 17, 184,
	3, 2, 2, 2, 19, 192, 3, 2, 2, 2, 21, 194, 3, 2, 2, 2, 23, 203, 3, 2, 2,
	2, 25, 211, 3, 2, 2, 2, 27, 217, 3, 2, 2, 2, 29, 225, 3, 2, 2, 2, 31, 234,
	3, 2, 2, 2, 33, 241, 3, 2, 2, 2, 35, 243, 3, 2, 2, 2, 37, 248, 3, 2, 2,
	2, 39, 251, 3, 2, 2, 2, 41, 258, 3, 2, 2, 2, 43, 263, 3, 2, 2, 2, 45, 269,
	3, 2, 2, 2, 47, 274, 3, 2, 2, 2, 49, 280, 3, 2, 2, 2, 51, 283, 3, 2, 2,
	2, 53, 298, 3, 2, 2, 2, 55, 312, 3, 2, 2, 2, 57, 326, 3, 2, 2, 2, 59, 340,
	3, 2, 2, 2, 61, 343, 3, 2, 2, 2, 63, 346, 3, 2, 2, 2, 65, 349, 3, 2, 2,
	2, 67, 352, 3, 2, 2, 2, 69, 355, 3, 2, 2, 2, 71, 357, 3, 2, 2, 2, 73, 359,
	3, 2, 2, 2, 75, 361, 3, 2, 2, 2, 77, 363, 3, 2, 2, 2, 79, 365, 3, 2, 2,
	2, 81, 367, 3, 2, 2, 2, 83, 372, 3, 2, 2, 2, 85, 378, 3, 2, 2, 2, 87, 403,
	3, 2, 2, 2, 89, 411, 3, 2, 2, 2, 91, 426, 3, 2, 2, 2, 93, 429, 3, 2, 2,
	2, 95, 434, 3, 2, 2, 2, 97, 462, 3, 2, 2, 2, 99, 474, 3, 2, 2, 2, 101,
	480, 3, 2, 2, 2, 103, 482, 3, 2, 2, 2, 105, 484, 3, 2, 2, 2, 107, 486,
	3, 2, 2, 2, 109, 488, 3, 2, 2, 2, 111, 490, 3, 2, 2, 2, 113, 492, 3, 2,
	2, 2, 115, 494, 3, 2, 2, 2, 117, 496, 3, 2, 2, 2, 119, 498, 3, 2, 2, 2,
	121, 500, 3, 2, 2, 2, 123, 502, 3, 2, 2, 2, 125, 504, 3, 2, 2, 2, 127,
	506, 3, 2, 2, 2, 129, 508, 3, 2, 2, 2, 131, 510, 3, 2, 2, 2, 133, 512,
	3, 2, 2, 2, 135, 514, 3, 2, 2, 2, 137, 516, 3, 2, 2, 2, 139, 518, 3, 2,
	2, 2, 141, 520, 3, 2, 2, 2, 143, 522, 3, 2, 2, 2, 145, 524, 3, 2, 2, 2,
	147, 526, 3, 2, 2, 2, 149, 528, 3, 2, 2, 2, 151, 530, 3, 2, 2, 2, 153,
	154, 7, 46, 2, 2, 154, 4, 3, 2, 2, 2, 155, 156, 5, 117, 59, 2, 156, 157,
	5, 127, 64, 2, 157, 158, 5, 137, 69, 2, 158, 159, 5, 109, 55, 2, 159, 160,
	5, 135, 68, 2, 160, 161, 5, 139, 70, 2, 161, 6, 3, 2, 2, 2, 162, 163, 5,
	117, 59, 2, 163, 164, 5, 127, 64, 2, 164, 165, 5, 139, 70, 2, 165, 166,
	5, 129, 65, 2, 166, 8, 3, 2, 2, 2, 167, 168, 5, 101, 51, 2, 168, 169, 5,
	137, 69, 2, 169, 10, 3, 2, 2, 2, 170, 171, 5, 101, 51, 2, 171, 172, 5,
	127, 64, 2, 172, 173, 5, 107, 54, 2, 173, 12, 3, 2, 2, 2, 174, 175, 5,
	105, 53, 2, 175, 176, 5, 101, 51, 2, 176, 177, 5, 137, 69, 2, 177, 178,
	5, 109, 55, 2, 178, 14, 3, 2, 2, 2, 179, 180, 5, 109, 55, 2, 180, 181,
	5, 123, 62, 2, 181, 182, 5, 137, 69, 2, 182, 183, 5, 109, 55, 2, 183, 16,
	3, 2, 2, 2, 184, 185, 5, 109, 55, 2, 185, 186, 5, 127, 64, 2, 186, 187,
	5, 107, 54, 2, 187, 18, 3, 2, 2, 2, 188, 189, 5, 109, 55, 2, 189, 190,
	5, 133, 67, 2, 190, 193, 3, 2, 2, 2, 191, 193, 7, 63, 2, 2, 192, 188, 3,
	2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 20, 3, 2, 2, 2, 194, 195, 5, 111, 56,
	2, 195, 196, 5, 135, 68, 2, 196, 197, 5, 129, 65, 2, 197, 198, 5, 125,
	63, 2, 198, 22, 3, 2, 2, 2, 199, 200, 5, 113, 57, 2, 200, 201, 5, 139,
	70, 2, 201, 204, 3, 2, 2, 2, 202, 204, 7, 64, 2, 2, 203, 199, 3, 2, 2,
	2, 203, 202, 3, 2, 2, 2, 204, 24, 3, 2, 2, 2, 205, 206, 5, 113, 57, 2,
	206, 207, 5, 139, 70, 2, 207, 208, 5, 109, 55, 2, 208, 212, 3, 2, 2, 2,
	209, 210, 7, 64, 2, 2, 210, 212, 7, 63, 2, 2, 211, 205, 3, 2, 2, 2, 211,
	209, 3, 2, 2, 2, 212, 26, 3, 2, 2, 2, 213, 214, 5, 123, 62, 2, 214, 215,
	5, 139, 70, 2, 215, 218, 3, 2, 2, 2, 216, 218, 7, 62, 2, 2, 217, 213, 3,
	2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 28, 3, 2, 2, 2, 219, 220, 5, 123, 62,
	2, 220, 221, 5, 139, 70, 2, 221, 222, 5, 109, 55, 2, 222, 226, 3, 2, 2,
	2, 223, 224, 7, 62, 2, 2, 224, 226, 7, 63, 2, 2, 225, 219, 3, 2, 2, 2,
	225, 223, 3, 2, 2, 2, 226, 30, 3, 2, 2, 2, 227, 228, 5, 127, 64, 2, 228,
	229, 5, 109, 55, 2, 229, 235, 3, 2, 2, 2, 230, 231, 7, 35, 2, 2, 231, 235,
	7, 63, 2, 2, 232, 233, 7, 62, 2, 2, 233, 235, 7, 64, 2, 2, 234, 227, 3,
	2, 2, 2, 234, 230, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 235, 32, 3, 2, 2,
	2, 236, 237, 5, 127, 64, 2, 237, 238, 5, 129, 65, 2, 238, 239, 5, 139,
	70, 2, 239, 242, 3, 2, 2, 2, 240, 242, 7, 35, 2, 2, 241, 236, 3, 2, 2,
	2, 241, 240, 3, 2, 2, 2, 242, 34, 3, 2, 2, 2, 243, 244, 5, 127, 64, 2,
	244, 245, 5, 141, 71, 2, 245, 246, 5, 123, 62, 2, 246, 247, 5, 123, 62,
	2, 247, 36, 3, 2, 2, 2, 248, 249, 5, 129, 65, 2, 249, 250, 5, 135, 68,
	2, 250, 38, 3, 2, 2, 2, 251, 252, 5, 137, 69, 2, 252, 253, 5, 109, 55,
	2, 253, 254, 5, 123, 62, 2, 254, 255, 5, 109, 55, 2, 255, 256, 5, 105,
	53, 2, 256, 257, 5, 139, 70, 2, 257, 40, 3, 2, 2, 2, 258, 259, 5, 139,
	70, 2, 259, 260, 5, 115, 58, 2, 260, 261, 5, 109, 55, 2, 261, 262, 5, 127,
	64, 2, 262, 42, 3, 2, 2, 2, 263, 264, 5, 145, 73, 2, 264, 265, 5, 115,
	58, 2, 265, 266, 5, 109, 55, 2, 266, 267, 5, 135, 68, 2, 267, 268, 5, 109,
	55, 2, 268, 44, 3, 2, 2, 2, 269, 270, 5, 145, 73, 2, 270, 271, 5, 115,
	58, 2, 271, 272, 5, 109, 55, 2, 272, 273, 5, 127, 64, 2, 273, 46, 3, 2,
	2, 2, 274, 275, 5, 113, 57, 2, 275, 276, 5, 135, 68, 2, 276, 277, 5, 129,
	65, 2, 277, 278, 5, 141, 71, 2, 278, 279, 5, 131, 66, 2, 279, 48, 3, 2,
	2, 2, 280, 281, 5, 103, 52, 2, 281, 282, 5, 149, 75, 2, 282, 50, 3, 2,
	2, 2, 283, 284, 5, 139, 70, 2, 284, 285, 5, 141, 71, 2, 285, 286, 5, 125,
	63, 2, 286, 287, 5, 103, 52, 2, 287, 288, 5, 123, 62, 2, 288, 289, 5, 117,
	59, 2, 289, 290, 5, 127, 64, 2, 290, 291, 5, 113, 57, 2, 291, 292, 5, 145,
	73, 2, 292, 293, 5, 117, 59, 2, 293, 294, 5, 127, 64, 2, 294, 295, 5, 107,
	54, 2, 295, 296, 5, 129, 65, 2, 296, 297, 5, 145, 73, 2, 297, 52, 3, 2,
	2, 2, 298, 299, 5, 115, 58, 2, 299, 300, 5, 129, 65, 2, 300, 301, 5, 131,
	66, 2, 301, 302, 5, 131, 66, 2, 302, 303, 5, 117, 59, 2, 303, 304, 5, 127,
	64, 2, 304, 305, 5, 113, 57, 2, 305, 306, 5, 145, 73, 2, 306, 307, 5, 117,
	59, 2, 307, 308, 5, 127, 64, 2, 308, 309, 5, 107, 54, 2, 309, 310, 5, 129,
	65, 2, 310, 311, 5, 145, 73, 2, 311, 54, 3, 2, 2, 2, 312, 313, 5, 137,
	69, 2, 313, 314, 5, 123, 62, 2, 314, 315, 5, 117, 59, 2, 315, 316, 5, 107,
	54, 2, 316, 317, 5, 117, 59, 2, 317, 318, 5, 127, 64, 2, 318, 319, 5, 113,
	57, 2, 319, 320, 5, 145, 73, 2, 320, 321, 5, 117, 59, 2, 321, 322, 5, 127,
	64, 2, 322, 323, 5, 107, 54, 2, 323, 324, 5, 129, 65, 2, 324, 325, 5, 145,
	73, 2, 325, 56, 3, 2, 2, 2, 326, 327, 5, 137, 69, 2, 327, 328, 5, 109,
	55, 2, 328, 329, 5, 137, 69, 2, 329, 330, 5, 137, 69, 2, 330, 331, 5, 117,
	59, 2, 331, 332, 5, 129, 65, 2, 332, 333, 5, 127, 64, 2, 333, 334, 5, 145,
	73, 2, 334, 335, 5, 117, 59, 2, 335, 336, 5, 127, 64, 2, 336, 337, 5, 107,
	54, 2, 337, 338, 5, 129, 65, 2, 338, 339, 5, 145, 73, 2, 339, 58, 3, 2,
	2, 2, 340, 341, 5, 107, 54, 2, 341, 342, 5, 107, 54, 2, 342, 60, 3, 2,
	2, 2, 343, 344, 5, 115, 58, 2, 344, 345, 5, 115, 58, 2, 345, 62, 3, 2,
	2, 2, 346, 347, 5, 125, 63, 2, 347, 348, 5, 117, 59, 2, 348, 64, 3, 2,
	2, 2, 349, 350, 5, 137, 69, 2, 350, 351, 5, 137, 69, 2, 351, 66, 3, 2,
	2, 2, 352, 353, 5, 125, 63, 2, 353, 354, 5, 137, 69, 2, 354, 68, 3, 2,
	2, 2, 355, 356, 7, 44, 2, 2, 356, 70, 3, 2, 2, 2, 357, 358, 7, 49, 2, 2,
	358, 72, 3, 2, 2, 2, 359, 360, 7, 39, 2, 2, 360, 74, 3, 2, 2, 2, 361, 362,
	7, 45, 2, 2, 362, 76, 3, 2, 2, 2, 363, 364, 7, 47, 2, 2, 364, 78, 3, 2,
	2, 2, 365, 366, 7, 48, 2, 2, 366, 80, 3, 2, 2, 2, 367, 368, 5, 139, 70,
	2, 368, 369, 5, 135, 68, 2, 369, 370, 5, 141, 71, 2, 370, 371, 5, 109,
	55, 2, 371, 82, 3, 2, 2, 2, 372, 373, 5, 111, 56, 2, 373, 374, 5, 101,
	51, 2, 374, 375, 5, 123, 62, 2, 375, 376, 5, 137, 69, 2, 376, 377, 5, 109,
	55, 2, 377, 84, 3, 2, 2, 2, 378, 401, 9, 2, 2, 2, 379, 381, 9, 3, 2, 2,
	380, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 382,
	383, 3, 2, 2, 2, 383, 387, 3, 2, 2, 2, 384, 386, 9, 4, 2, 2, 385, 384,
	3, 2, 2, 2, 386, 389, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 388, 3, 2,
	2, 2, 388, 402, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 390, 392, 9, 4, 2, 2,
	391, 390, 3, 2, 2, 2, 392, 395, 3, 2, 2, 2, 393, 391, 3, 2, 2, 2, 393,
	394, 3, 2, 2, 2, 394, 397, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 396, 398,
	9, 3, 2, 2, 397, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 397, 3, 2,
	2, 2, 399, 400, 3, 2, 2, 2, 400, 402, 3, 2, 2, 2, 401, 380, 3, 2, 2, 2,
	401, 393, 3, 2, 2, 2, 402, 86, 3, 2, 2, 2, 403, 404, 7, 48, 2, 2, 404,
	408, 9, 5, 2, 2, 405, 407, 9, 6, 2, 2, 406, 405, 3, 2, 2, 2, 407, 410,
	3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 88, 3, 2,
	2, 2, 410, 408, 3, 2, 2, 2, 411, 415, 9, 7, 2, 2, 412, 414, 9, 8, 2, 2,
	413, 412, 3, 2, 2, 2, 414, 417, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 415,
	416, 3, 2, 2, 2, 416, 90, 3, 2, 2, 2, 417, 415, 3, 2, 2, 2, 418, 427, 7,
	50, 2, 2, 419, 423, 9, 9, 2, 2, 420, 422, 9, 4, 2, 2, 421, 420, 3, 2, 2,
	2, 422, 425, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424,
	427, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 418, 3, 2, 2, 2, 426, 419,
	3, 2, 2, 2, 427, 92, 3, 2, 2, 2, 428, 430, 9, 10, 2, 2, 429, 428, 3, 2,
	2, 2, 429, 430, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 432, 5, 91, 46,
	2, 432, 94, 3, 2, 2, 2, 433, 435, 9, 10, 2, 2, 434, 433, 3, 2, 2, 2, 434,
	435, 3, 2, 2, 2, 435, 460, 3, 2, 2, 2, 436, 438, 5, 91, 46, 2, 437, 436,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2,
	2, 2, 440, 441, 3, 2, 2, 2, 441, 443, 5, 79, 40, 2, 442, 444, 5, 91, 46,
	2, 443, 442, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 443, 3, 2, 2, 2, 445,
	446, 3, 2, 2, 2, 446, 461, 3, 2, 2, 2, 447, 449, 5, 91, 46, 2, 448, 447,
	3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 448, 3, 2, 2, 2, 450, 451, 3, 2,
	2, 2, 451, 452, 3, 2, 2, 2, 452, 453, 5, 79, 40, 2, 453, 461, 3, 2, 2,
	2, 454, 456, 5, 79, 40, 2, 455, 457, 5, 91, 46, 2, 456, 455, 3, 2, 2, 2,
	457, 458, 3, 2, 2, 2, 458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459,
	461, 3, 2, 2, 2, 460, 437, 3, 2, 2, 2, 460, 448, 3, 2, 2, 2, 460, 454,
	3, 2, 2, 2, 461, 96, 3, 2, 2, 2, 462, 468, 7, 41, 2, 2, 463, 467, 10, 11,
	2, 2, 464, 465, 7, 41, 2, 2, 465, 467, 7, 41, 2, 2, 466, 463, 3, 2, 2,
	2, 466, 464, 3, 2, 2, 2, 467, 470, 3, 2, 2, 2, 468, 466, 3, 2, 2, 2, 468,
	469, 3, 2, 2, 2, 469, 471, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 471, 472,
	7, 41, 2, 2, 472, 98, 3, 2, 2, 2, 473, 475, 9, 12, 2, 2, 474, 473, 3, 2,
	2, 2, 475, 476, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2,
	477, 478, 3, 2, 2, 2, 478, 479, 8, 50, 2, 2, 479, 100, 3, 2, 2, 2, 480,
	481, 9, 13, 2, 2, 481, 102, 3, 2, 2, 2, 482, 483, 9, 14, 2, 2, 483, 104,
	3, 2, 2, 2, 484, 485, 9, 15, 2, 2, 485, 106, 3, 2, 2, 2, 486, 487, 9, 16,
	2, 2, 487, 108, 3, 2, 2, 2, 488, 489, 9, 17, 2, 2, 489, 110, 3, 2, 2, 2,
	490, 491, 9, 18, 2, 2, 491, 112, 3, 2, 2, 2, 492, 493, 9, 19, 2, 2, 493,
	114, 3, 2, 2, 2, 494, 495, 9, 20, 2, 2, 495, 116, 3, 2, 2, 2, 496, 497,
	9, 21, 2, 2, 497, 118, 3, 2, 2, 2, 498, 499, 9, 22, 2, 2, 499, 120, 3,
	2, 2, 2, 500, 501, 9, 23, 2, 2, 501, 122, 3, 2, 2, 2, 502, 503, 9, 24,
	2, 2, 503, 124, 3, 2, 2, 2, 504, 505, 9, 25, 2, 2, 505, 126, 3, 2, 2, 2,
	506, 507, 9, 26, 2, 2, 507, 128, 3, 2, 2, 2, 508, 509, 9, 27, 2, 2, 509,
	130, 3, 2, 2, 2, 510, 511, 9, 28, 2, 2, 511, 132, 3, 2, 2, 2, 512, 513,
	9, 29, 2, 2, 513, 134, 3, 2, 2, 2, 514, 515, 9, 30, 2, 2, 515, 136, 3,
	2, 2, 2, 516, 517, 9, 31, 2, 2, 517, 138, 3, 2, 2, 2, 518, 519, 9, 32,
	2, 2, 519, 140, 3, 2, 2, 2, 520, 521, 9, 33, 2, 2, 521, 142, 3, 2, 2, 2,
	522, 523, 9, 34, 2, 2, 523, 144, 3, 2, 2, 2, 524, 525, 9, 35, 2, 2, 525,
	146, 3, 2, 2, 2, 526, 527, 9, 36, 2, 2, 527, 148, 3, 2, 2, 2, 528, 529,
	9, 37, 2, 2, 529, 150, 3, 2, 2, 2, 530, 531, 9, 38, 2, 2, 531, 152, 3,
	2, 2, 2, 29, 2, 192, 203, 211, 217, 225, 234, 241, 382, 387, 393, 399,
	401, 408, 415, 423, 426, 429, 434, 439, 445, 450, 458, 460, 466, 468, 476,
	3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'*'",
	"'/'", "'%'", "'+'", "'-'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "INSERT", "INTO", "AS", "AND", "CASE", "ELSE", "END", "EQ", "FROM",
	"GT", "GTE", "LT", "LTE", "NE", "NOT", "NULL", "OR", "SELECT", "THEN",
	"WHERE", "WHEN", "GROUP", "BY", "TUMBLINGWINDOW", "HOPPINGWINDOW", "SLIDINGWINDOW",
	"SESSIONWINDOW", "DD", "HH", "MI", "SS", "MS", "MUL", "DIV", "MOD", "ADD",
	"SUB", "DOT", "TRUE", "FALSE", "ENTITYNAME", "PROPERTYNAME", "TARGETENTITY",
	"NUMBER", "INTEGER", "FLOAT", "STRING", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "INSERT", "INTO", "AS", "AND", "CASE", "ELSE", "END", "EQ", "FROM",
	"GT", "GTE", "LT", "LTE", "NE", "NOT", "NULL", "OR", "SELECT", "THEN",
	"WHERE", "WHEN", "GROUP", "BY", "TUMBLINGWINDOW", "HOPPINGWINDOW", "SLIDINGWINDOW",
	"SESSIONWINDOW", "DD", "HH", "MI", "SS", "MS", "MUL", "DIV", "MOD", "ADD",
	"SUB", "DOT", "TRUE", "FALSE", "ENTITYNAME", "PROPERTYNAME", "TARGETENTITY",
	"NUMBER", "INTEGER", "FLOAT", "STRING", "WHITESPACE", "A", "B", "C", "D",
	"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	"T", "U", "V", "W", "X", "Y", "Z",
}

type TQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTQLLexer(input antlr.CharStream) *TQLLexer {

	l := new(TQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TQLLexer tokens.
const (
	TQLLexerT__0           = 1
	TQLLexerINSERT         = 2
	TQLLexerINTO           = 3
	TQLLexerAS             = 4
	TQLLexerAND            = 5
	TQLLexerCASE           = 6
	TQLLexerELSE           = 7
	TQLLexerEND            = 8
	TQLLexerEQ             = 9
	TQLLexerFROM           = 10
	TQLLexerGT             = 11
	TQLLexerGTE            = 12
	TQLLexerLT             = 13
	TQLLexerLTE            = 14
	TQLLexerNE             = 15
	TQLLexerNOT            = 16
	TQLLexerNULL           = 17
	TQLLexerOR             = 18
	TQLLexerSELECT         = 19
	TQLLexerTHEN           = 20
	TQLLexerWHERE          = 21
	TQLLexerWHEN           = 22
	TQLLexerGROUP          = 23
	TQLLexerBY             = 24
	TQLLexerTUMBLINGWINDOW = 25
	TQLLexerHOPPINGWINDOW  = 26
	TQLLexerSLIDINGWINDOW  = 27
	TQLLexerSESSIONWINDOW  = 28
	TQLLexerDD             = 29
	TQLLexerHH             = 30
	TQLLexerMI             = 31
	TQLLexerSS             = 32
	TQLLexerMS             = 33
	TQLLexerMUL            = 34
	TQLLexerDIV            = 35
	TQLLexerMOD            = 36
	TQLLexerADD            = 37
	TQLLexerSUB            = 38
	TQLLexerDOT            = 39
	TQLLexerTRUE           = 40
	TQLLexerFALSE          = 41
	TQLLexerENTITYNAME     = 42
	TQLLexerPROPERTYNAME   = 43
	TQLLexerTARGETENTITY   = 44
	TQLLexerNUMBER         = 45
	TQLLexerINTEGER        = 46
	TQLLexerFLOAT          = 47
	TQLLexerSTRING         = 48
	TQLLexerWHITESPACE     = 49
)
