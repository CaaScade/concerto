// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 694,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 5, 2, 77,
	10, 2, 3, 2, 7, 2, 80, 10, 2, 12, 2, 14, 2, 83, 11, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 6, 3, 89, 10, 3, 13, 3, 14, 3, 90, 3, 3, 3, 3, 6, 3, 95, 10, 3, 13,
	3, 14, 3, 96, 3, 4, 3, 4, 3, 4, 6, 4, 102, 10, 4, 13, 4, 14, 4, 103, 3,
	4, 7, 4, 107, 10, 4, 12, 4, 14, 4, 110, 11, 4, 3, 5, 7, 5, 113, 10, 5,
	12, 5, 14, 5, 116, 11, 5, 3, 5, 3, 5, 3, 5, 7, 5, 121, 10, 5, 12, 5, 14,
	5, 124, 11, 5, 3, 5, 3, 5, 7, 5, 128, 10, 5, 12, 5, 14, 5, 131, 11, 5,
	3, 5, 3, 5, 7, 5, 135, 10, 5, 12, 5, 14, 5, 138, 11, 5, 3, 5, 3, 5, 5,
	5, 142, 10, 5, 3, 6, 7, 6, 145, 10, 6, 12, 6, 14, 6, 148, 11, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 155, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 160,
	10, 8, 12, 8, 14, 8, 163, 11, 8, 3, 8, 5, 8, 166, 10, 8, 3, 8, 3, 8, 7,
	8, 170, 10, 8, 12, 8, 14, 8, 173, 11, 8, 3, 9, 3, 9, 7, 9, 177, 10, 9,
	12, 9, 14, 9, 180, 11, 9, 3, 9, 3, 9, 3, 9, 7, 9, 185, 10, 9, 12, 9, 14,
	9, 188, 11, 9, 3, 9, 5, 9, 191, 10, 9, 3, 9, 3, 9, 7, 9, 195, 10, 9, 12,
	9, 14, 9, 198, 11, 9, 3, 10, 7, 10, 201, 10, 10, 12, 10, 14, 10, 204, 11,
	10, 3, 10, 3, 10, 7, 10, 208, 10, 10, 12, 10, 14, 10, 211, 11, 10, 3, 10,
	3, 10, 3, 10, 7, 10, 216, 10, 10, 12, 10, 14, 10, 219, 11, 10, 7, 10, 221,
	10, 10, 12, 10, 14, 10, 224, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	5, 11, 231, 10, 11, 3, 12, 3, 12, 6, 12, 235, 10, 12, 13, 12, 14, 12, 236,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 248,
	10, 14, 12, 14, 14, 14, 251, 11, 14, 3, 14, 3, 14, 7, 14, 255, 10, 14,
	12, 14, 14, 14, 258, 11, 14, 3, 14, 7, 14, 261, 10, 14, 12, 14, 14, 14,
	264, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 270, 10, 15, 12, 15, 14,
	15, 273, 11, 15, 3, 15, 3, 15, 7, 15, 277, 10, 15, 12, 15, 14, 15, 280,
	11, 15, 3, 15, 3, 15, 7, 15, 284, 10, 15, 12, 15, 14, 15, 287, 11, 15,
	7, 15, 289, 10, 15, 12, 15, 14, 15, 292, 11, 15, 5, 15, 294, 10, 15, 3,
	15, 7, 15, 297, 10, 15, 12, 15, 14, 15, 300, 11, 15, 3, 15, 5, 15, 303,
	10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 7, 15, 315, 10, 15, 12, 15, 14, 15, 318, 11, 15, 3, 16, 3, 16, 7,
	16, 322, 10, 16, 12, 16, 14, 16, 325, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	7, 17, 331, 10, 17, 12, 17, 14, 17, 334, 11, 17, 3, 17, 3, 17, 5, 17, 338,
	10, 17, 3, 17, 7, 17, 341, 10, 17, 12, 17, 14, 17, 344, 11, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 352, 10, 18, 12, 18, 14, 18,
	355, 11, 18, 3, 18, 3, 18, 7, 18, 359, 10, 18, 12, 18, 14, 18, 362, 11,
	18, 3, 18, 3, 18, 3, 18, 5, 18, 367, 10, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 3, 21, 5, 21, 375, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 6, 23, 383, 10, 23, 13, 23, 14, 23, 384, 3, 23, 3, 23, 7, 23, 389,
	10, 23, 12, 23, 14, 23, 392, 11, 23, 3, 23, 3, 23, 5, 23, 396, 10, 23,
	3, 23, 3, 23, 6, 23, 400, 10, 23, 13, 23, 14, 23, 401, 3, 23, 3, 23, 7,
	23, 406, 10, 23, 12, 23, 14, 23, 409, 11, 23, 3, 23, 3, 23, 7, 23, 413,
	10, 23, 12, 23, 14, 23, 416, 11, 23, 3, 23, 5, 23, 419, 10, 23, 3, 24,
	3, 24, 6, 24, 423, 10, 24, 13, 24, 14, 24, 424, 3, 25, 3, 25, 5, 25, 429,
	10, 25, 3, 26, 3, 26, 6, 26, 433, 10, 26, 13, 26, 14, 26, 434, 3, 26, 3,
	26, 7, 26, 439, 10, 26, 12, 26, 14, 26, 442, 11, 26, 3, 26, 3, 26, 7, 26,
	446, 10, 26, 12, 26, 14, 26, 449, 11, 26, 3, 26, 3, 26, 7, 26, 453, 10,
	26, 12, 26, 14, 26, 456, 11, 26, 6, 26, 458, 10, 26, 13, 26, 14, 26, 459,
	5, 26, 462, 10, 26, 3, 26, 7, 26, 465, 10, 26, 12, 26, 14, 26, 468, 11,
	26, 3, 26, 5, 26, 471, 10, 26, 3, 26, 7, 26, 474, 10, 26, 12, 26, 14, 26,
	477, 11, 26, 3, 26, 3, 26, 7, 26, 481, 10, 26, 12, 26, 14, 26, 484, 11,
	26, 3, 26, 7, 26, 487, 10, 26, 12, 26, 14, 26, 490, 11, 26, 3, 26, 7, 26,
	493, 10, 26, 12, 26, 14, 26, 496, 11, 26, 3, 26, 3, 26, 3, 27, 7, 27, 501,
	10, 27, 12, 27, 14, 27, 504, 11, 27, 3, 27, 3, 27, 7, 27, 508, 10, 27,
	12, 27, 14, 27, 511, 11, 27, 3, 27, 3, 27, 6, 27, 515, 10, 27, 13, 27,
	14, 27, 516, 3, 28, 3, 28, 6, 28, 521, 10, 28, 13, 28, 14, 28, 522, 3,
	28, 3, 28, 7, 28, 527, 10, 28, 12, 28, 14, 28, 530, 11, 28, 3, 28, 5, 28,
	533, 10, 28, 3, 28, 7, 28, 536, 10, 28, 12, 28, 14, 28, 539, 11, 28, 3,
	28, 3, 28, 7, 28, 543, 10, 28, 12, 28, 14, 28, 546, 11, 28, 3, 28, 3, 28,
	6, 28, 550, 10, 28, 13, 28, 14, 28, 551, 7, 28, 554, 10, 28, 12, 28, 14,
	28, 557, 11, 28, 3, 28, 7, 28, 560, 10, 28, 12, 28, 14, 28, 563, 11, 28,
	3, 28, 3, 28, 3, 29, 7, 29, 568, 10, 29, 12, 29, 14, 29, 571, 11, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 7, 30, 578, 10, 30, 12, 30, 14, 30, 581,
	11, 30, 3, 30, 3, 30, 3, 30, 7, 30, 586, 10, 30, 12, 30, 14, 30, 589, 11,
	30, 5, 30, 591, 10, 30, 3, 30, 7, 30, 594, 10, 30, 12, 30, 14, 30, 597,
	11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32, 606, 10,
	32, 12, 32, 14, 32, 609, 11, 32, 3, 32, 3, 32, 3, 32, 7, 32, 614, 10, 32,
	12, 32, 14, 32, 617, 11, 32, 5, 32, 619, 10, 32, 3, 32, 3, 32, 3, 33, 7,
	33, 624, 10, 33, 12, 33, 14, 33, 627, 11, 33, 3, 33, 3, 33, 7, 33, 631,
	10, 33, 12, 33, 14, 33, 634, 11, 33, 3, 33, 3, 33, 7, 33, 638, 10, 33,
	12, 33, 14, 33, 641, 11, 33, 3, 33, 7, 33, 644, 10, 33, 12, 33, 14, 33,
	647, 11, 33, 3, 33, 3, 33, 7, 33, 651, 10, 33, 12, 33, 14, 33, 654, 11,
	33, 5, 33, 656, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34,
	664, 10, 34, 3, 35, 3, 35, 7, 35, 668, 10, 35, 12, 35, 14, 35, 671, 11,
	35, 3, 35, 3, 35, 3, 35, 5, 35, 676, 10, 35, 3, 35, 7, 35, 679, 10, 35,
	12, 35, 14, 35, 682, 11, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 2, 4, 26, 28, 38, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 2, 5, 3, 2, 44, 45, 3,
	2, 22, 41, 4, 2, 43, 43, 48, 51, 2, 763, 2, 74, 3, 2, 2, 2, 4, 86, 3, 2,
	2, 2, 6, 98, 3, 2, 2, 2, 8, 141, 3, 2, 2, 2, 10, 146, 3, 2, 2, 2, 12, 154,
	3, 2, 2, 2, 14, 156, 3, 2, 2, 2, 16, 174, 3, 2, 2, 2, 18, 202, 3, 2, 2,
	2, 20, 230, 3, 2, 2, 2, 22, 232, 3, 2, 2, 2, 24, 240, 3, 2, 2, 2, 26, 242,
	3, 2, 2, 2, 28, 302, 3, 2, 2, 2, 30, 319, 3, 2, 2, 2, 32, 328, 3, 2, 2,
	2, 34, 366, 3, 2, 2, 2, 36, 368, 3, 2, 2, 2, 38, 370, 3, 2, 2, 2, 40, 374,
	3, 2, 2, 2, 42, 376, 3, 2, 2, 2, 44, 418, 3, 2, 2, 2, 46, 420, 3, 2, 2,
	2, 48, 428, 3, 2, 2, 2, 50, 430, 3, 2, 2, 2, 52, 502, 3, 2, 2, 2, 54, 518,
	3, 2, 2, 2, 56, 569, 3, 2, 2, 2, 58, 574, 3, 2, 2, 2, 60, 600, 3, 2, 2,
	2, 62, 602, 3, 2, 2, 2, 64, 655, 3, 2, 2, 2, 66, 663, 3, 2, 2, 2, 68, 665,
	3, 2, 2, 2, 70, 686, 3, 2, 2, 2, 72, 691, 3, 2, 2, 2, 74, 76, 5, 4, 3,
	2, 75, 77, 5, 6, 4, 2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 81,
	3, 2, 2, 2, 78, 80, 5, 12, 7, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2,
	81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3,
	2, 2, 2, 84, 85, 7, 2, 2, 3, 85, 3, 3, 2, 2, 2, 86, 88, 7, 3, 2, 2, 87,
	89, 7, 44, 2, 2, 88, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 88, 3, 2,
	2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 7, 42, 2, 2, 93,
	95, 5, 10, 6, 2, 94, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2,
	2, 2, 96, 97, 3, 2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 99, 7, 4, 2, 2, 99, 101,
	5, 10, 6, 2, 100, 102, 5, 8, 5, 2, 101, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 108, 3, 2, 2, 2,
	105, 107, 5, 10, 6, 2, 106, 105, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 7, 3, 2, 2, 2, 110, 108, 3,
	2, 2, 2, 111, 113, 7, 44, 2, 2, 112, 111, 3, 2, 2, 2, 113, 116, 3, 2, 2,
	2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117, 3, 2, 2, 2, 116,
	114, 3, 2, 2, 2, 117, 118, 7, 42, 2, 2, 118, 142, 5, 10, 6, 2, 119, 121,
	7, 44, 2, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2,
	2, 2, 122, 123, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2,
	125, 129, 7, 42, 2, 2, 126, 128, 7, 44, 2, 2, 127, 126, 3, 2, 2, 2, 128,
	131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132,
	3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 136, 7, 5, 2, 2, 133, 135, 7, 44,
	2, 2, 134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2,
	136, 137, 3, 2, 2, 2, 137, 139, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139,
	140, 7, 43, 2, 2, 140, 142, 5, 10, 6, 2, 141, 114, 3, 2, 2, 2, 141, 122,
	3, 2, 2, 2, 142, 9, 3, 2, 2, 2, 143, 145, 7, 44, 2, 2, 144, 143, 3, 2,
	2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2,
	147, 149, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 45, 2, 2, 150,
	11, 3, 2, 2, 2, 151, 155, 5, 46, 24, 2, 152, 155, 5, 14, 8, 2, 153, 155,
	5, 16, 9, 2, 154, 151, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2,
	2, 2, 155, 13, 3, 2, 2, 2, 156, 157, 7, 6, 2, 2, 157, 161, 5, 56, 29, 2,
	158, 160, 7, 44, 2, 2, 159, 158, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161,
	159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161,
	3, 2, 2, 2, 164, 166, 7, 45, 2, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 167, 3, 2, 2, 2, 167, 171, 5, 18, 10, 2, 168, 170, 9, 2, 2,
	2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171,
	172, 3, 2, 2, 2, 172, 15, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 174, 178, 7,
	6, 2, 2, 175, 177, 7, 44, 2, 2, 176, 175, 3, 2, 2, 2, 177, 180, 3, 2, 2,
	2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180,
	178, 3, 2, 2, 2, 181, 182, 7, 42, 2, 2, 182, 186, 5, 56, 29, 2, 183, 185,
	7, 44, 2, 2, 184, 183, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2,
	2, 2, 186, 187, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2,
	189, 191, 7, 45, 2, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	192, 3, 2, 2, 2, 192, 196, 5, 18, 10, 2, 193, 195, 9, 2, 2, 2, 194, 193,
	3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2,
	2, 2, 197, 17, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199, 201, 7, 44, 2, 2,
	200, 199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202,
	203, 3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 209,
	7, 7, 2, 2, 206, 208, 9, 2, 2, 2, 207, 206, 3, 2, 2, 2, 208, 211, 3, 2,
	2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 222, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 212, 213, 5, 20, 11, 2, 213, 217, 5, 10, 6, 2, 214,
	216, 9, 2, 2, 2, 215, 214, 3, 2, 2, 2, 216, 219, 3, 2, 2, 2, 217, 215,
	3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 221, 3, 2, 2, 2, 219, 217, 3, 2,
	2, 2, 220, 212, 3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2,
	222, 223, 3, 2, 2, 2, 223, 225, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225,
	226, 7, 8, 2, 2, 226, 19, 3, 2, 2, 2, 227, 231, 5, 24, 13, 2, 228, 231,
	5, 26, 14, 2, 229, 231, 5, 22, 12, 2, 230, 227, 3, 2, 2, 2, 230, 228, 3,
	2, 2, 2, 230, 229, 3, 2, 2, 2, 231, 21, 3, 2, 2, 2, 232, 234, 7, 9, 2,
	2, 233, 235, 7, 44, 2, 2, 234, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236,
	234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239,
	5, 26, 14, 2, 239, 23, 3, 2, 2, 2, 240, 241, 5, 44, 23, 2, 241, 25, 3,
	2, 2, 2, 242, 243, 8, 14, 1, 2, 243, 244, 5, 28, 15, 2, 244, 262, 3, 2,
	2, 2, 245, 249, 12, 3, 2, 2, 246, 248, 7, 44, 2, 2, 247, 246, 3, 2, 2,
	2, 248, 251, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250,
	252, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 252, 256, 9, 3, 2, 2, 253, 255,
	7, 44, 2, 2, 254, 253, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2,
	2, 2, 256, 257, 3, 2, 2, 2, 257, 259, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2,
	259, 261, 5, 26, 14, 4, 260, 245, 3, 2, 2, 2, 261, 264, 3, 2, 2, 2, 262,
	260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 27, 3, 2, 2, 2, 264, 262, 3,
	2, 2, 2, 265, 266, 8, 15, 1, 2, 266, 303, 5, 34, 18, 2, 267, 271, 7, 12,
	2, 2, 268, 270, 7, 44, 2, 2, 269, 268, 3, 2, 2, 2, 270, 273, 3, 2, 2, 2,
	271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 293, 3, 2, 2, 2, 273,
	271, 3, 2, 2, 2, 274, 290, 5, 28, 15, 2, 275, 277, 7, 44, 2, 2, 276, 275,
	3, 2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2,
	2, 2, 279, 281, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 281, 285, 5, 30, 16,
	2, 282, 284, 7, 44, 2, 2, 283, 282, 3, 2, 2, 2, 284, 287, 3, 2, 2, 2, 285,
	283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 289, 3, 2, 2, 2, 287, 285,
	3, 2, 2, 2, 288, 278, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3, 2,
	2, 2, 290, 291, 3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2,
	293, 274, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 298, 3, 2, 2, 2, 295,
	297, 7, 44, 2, 2, 296, 295, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296,
	3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 301, 3, 2, 2, 2, 300, 298, 3, 2,
	2, 2, 301, 303, 7, 13, 2, 2, 302, 265, 3, 2, 2, 2, 302, 267, 3, 2, 2, 2,
	303, 316, 3, 2, 2, 2, 304, 305, 12, 7, 2, 2, 305, 306, 7, 10, 2, 2, 306,
	315, 7, 42, 2, 2, 307, 308, 12, 6, 2, 2, 308, 309, 7, 10, 2, 2, 309, 315,
	5, 58, 30, 2, 310, 311, 12, 5, 2, 2, 311, 315, 7, 11, 2, 2, 312, 313, 12,
	4, 2, 2, 313, 315, 5, 32, 17, 2, 314, 304, 3, 2, 2, 2, 314, 307, 3, 2,
	2, 2, 314, 310, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 315, 318, 3, 2, 2, 2,
	316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 29, 3, 2, 2, 2, 318, 316,
	3, 2, 2, 2, 319, 323, 7, 14, 2, 2, 320, 322, 7, 44, 2, 2, 321, 320, 3,
	2, 2, 2, 322, 325, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 323, 324, 3, 2, 2,
	2, 324, 326, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 326, 327, 5, 28, 15, 2,
	327, 31, 3, 2, 2, 2, 328, 332, 7, 12, 2, 2, 329, 331, 7, 44, 2, 2, 330,
	329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332, 333,
	3, 2, 2, 2, 333, 337, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 335, 338, 5, 28,
	15, 2, 336, 338, 5, 72, 37, 2, 337, 335, 3, 2, 2, 2, 337, 336, 3, 2, 2,
	2, 338, 342, 3, 2, 2, 2, 339, 341, 7, 44, 2, 2, 340, 339, 3, 2, 2, 2, 341,
	344, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 345,
	3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 345, 346, 7, 13, 2, 2, 346, 33, 3, 2,
	2, 2, 347, 367, 5, 58, 30, 2, 348, 367, 5, 40, 21, 2, 349, 353, 7, 15,
	2, 2, 350, 352, 7, 44, 2, 2, 351, 350, 3, 2, 2, 2, 352, 355, 3, 2, 2, 2,
	353, 351, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 356, 3, 2, 2, 2, 355,
	353, 3, 2, 2, 2, 356, 360, 5, 26, 14, 2, 357, 359, 7, 44, 2, 2, 358, 357,
	3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2,
	2, 2, 361, 363, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 364, 7, 16, 2, 2,
	364, 367, 3, 2, 2, 2, 365, 367, 5, 36, 19, 2, 366, 347, 3, 2, 2, 2, 366,
	348, 3, 2, 2, 2, 366, 349, 3, 2, 2, 2, 366, 365, 3, 2, 2, 2, 367, 35, 3,
	2, 2, 2, 368, 369, 5, 38, 20, 2, 369, 37, 3, 2, 2, 2, 370, 371, 9, 4, 2,
	2, 371, 39, 3, 2, 2, 2, 372, 375, 7, 42, 2, 2, 373, 375, 5, 42, 22, 2,
	374, 372, 3, 2, 2, 2, 374, 373, 3, 2, 2, 2, 375, 41, 3, 2, 2, 2, 376, 377,
	7, 42, 2, 2, 377, 378, 7, 10, 2, 2, 378, 379, 7, 42, 2, 2, 379, 43, 3,
	2, 2, 2, 380, 382, 7, 17, 2, 2, 381, 383, 7, 44, 2, 2, 382, 381, 3, 2,
	2, 2, 383, 384, 3, 2, 2, 2, 384, 382, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2,
	385, 386, 3, 2, 2, 2, 386, 390, 7, 42, 2, 2, 387, 389, 7, 44, 2, 2, 388,
	387, 3, 2, 2, 2, 389, 392, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 391,
	3, 2, 2, 2, 391, 395, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 393, 396, 7, 42,
	2, 2, 394, 396, 5, 58, 30, 2, 395, 393, 3, 2, 2, 2, 395, 394, 3, 2, 2,
	2, 396, 419, 3, 2, 2, 2, 397, 399, 7, 17, 2, 2, 398, 400, 7, 44, 2, 2,
	399, 398, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 401,
	402, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 407, 7, 42, 2, 2, 404, 406,
	7, 44, 2, 2, 405, 404, 3, 2, 2, 2, 406, 409, 3, 2, 2, 2, 407, 405, 3, 2,
	2, 2, 407, 408, 3, 2, 2, 2, 408, 410, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2,
	410, 414, 7, 38, 2, 2, 411, 413, 7, 44, 2, 2, 412, 411, 3, 2, 2, 2, 413,
	416, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 417,
	3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 417, 419, 5, 26, 14, 2, 418, 380, 3,
	2, 2, 2, 418, 397, 3, 2, 2, 2, 419, 45, 3, 2, 2, 2, 420, 422, 5, 48, 25,
	2, 421, 423, 5, 10, 6, 2, 422, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424,
	422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 47, 3, 2, 2, 2, 426, 429, 5,
	50, 26, 2, 427, 429, 5, 54, 28, 2, 428, 426, 3, 2, 2, 2, 428, 427, 3, 2,
	2, 2, 429, 49, 3, 2, 2, 2, 430, 432, 7, 18, 2, 2, 431, 433, 7, 44, 2, 2,
	432, 431, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 434,
	435, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 440, 7, 42, 2, 2, 437, 439,
	7, 44, 2, 2, 438, 437, 3, 2, 2, 2, 439, 442, 3, 2, 2, 2, 440, 438, 3, 2,
	2, 2, 440, 441, 3, 2, 2, 2, 441, 461, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2,
	443, 457, 7, 19, 2, 2, 444, 446, 7, 44, 2, 2, 445, 444, 3, 2, 2, 2, 446,
	449, 3, 2, 2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 450,
	3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 450, 454, 7, 42, 2, 2, 451, 453, 7, 44,
	2, 2, 452, 451, 3, 2, 2, 2, 453, 456, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2,
	454, 455, 3, 2, 2, 2, 455, 458, 3, 2, 2, 2, 456, 454, 3, 2, 2, 2, 457,
	447, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 457, 3, 2, 2, 2, 459, 460,
	3, 2, 2, 2, 460, 462, 3, 2, 2, 2, 461, 443, 3, 2, 2, 2, 461, 462, 3, 2,
	2, 2, 462, 466, 3, 2, 2, 2, 463, 465, 7, 44, 2, 2, 464, 463, 3, 2, 2, 2,
	465, 468, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 466, 467, 3, 2, 2, 2, 467,
	470, 3, 2, 2, 2, 468, 466, 3, 2, 2, 2, 469, 471, 7, 45, 2, 2, 470, 469,
	3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 475, 3, 2, 2, 2, 472, 474, 7, 44,
	2, 2, 473, 472, 3, 2, 2, 2, 474, 477, 3, 2, 2, 2, 475, 473, 3, 2, 2, 2,
	475, 476, 3, 2, 2, 2, 476, 478, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 478,
	482, 7, 7, 2, 2, 479, 481, 9, 2, 2, 2, 480, 479, 3, 2, 2, 2, 481, 484,
	3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 488, 3, 2,
	2, 2, 484, 482, 3, 2, 2, 2, 485, 487, 5, 52, 27, 2, 486, 485, 3, 2, 2,
	2, 487, 490, 3, 2, 2, 2, 488, 486, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489,
	494, 3, 2, 2, 2, 490, 488, 3, 2, 2, 2, 491, 493, 9, 2, 2, 2, 492, 491,
	3, 2, 2, 2, 493, 496, 3, 2, 2, 2, 494, 492, 3, 2, 2, 2, 494, 495, 3, 2,
	2, 2, 495, 497, 3, 2, 2, 2, 496, 494, 3, 2, 2, 2, 497, 498, 7, 8, 2, 2,
	498, 51, 3, 2, 2, 2, 499, 501, 7, 44, 2, 2, 500, 499, 3, 2, 2, 2, 501,
	504, 3, 2, 2, 2, 502, 500, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 505,
	3, 2, 2, 2, 504, 502, 3, 2, 2, 2, 505, 509, 7, 42, 2, 2, 506, 508, 7, 44,
	2, 2, 507, 506, 3, 2, 2, 2, 508, 511, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2,
	509, 510, 3, 2, 2, 2, 510, 512, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 512,
	514, 5, 66, 34, 2, 513, 515, 5, 10, 6, 2, 514, 513, 3, 2, 2, 2, 515, 516,
	3, 2, 2, 2, 516, 514, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 53, 3, 2,
	2, 2, 518, 520, 7, 20, 2, 2, 519, 521, 7, 44, 2, 2, 520, 519, 3, 2, 2,
	2, 521, 522, 3, 2, 2, 2, 522, 520, 3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523,
	524, 3, 2, 2, 2, 524, 528, 7, 42, 2, 2, 525, 527, 7, 44, 2, 2, 526, 525,
	3, 2, 2, 2, 527, 530, 3, 2, 2, 2, 528, 526, 3, 2, 2, 2, 528, 529, 3, 2,
	2, 2, 529, 532, 3, 2, 2, 2, 530, 528, 3, 2, 2, 2, 531, 533, 7, 45, 2, 2,
	532, 531, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 537, 3, 2, 2, 2, 534,
	536, 7, 44, 2, 2, 535, 534, 3, 2, 2, 2, 536, 539, 3, 2, 2, 2, 537, 535,
	3, 2, 2, 2, 537, 538, 3, 2, 2, 2, 538, 540, 3, 2, 2, 2, 539, 537, 3, 2,
	2, 2, 540, 544, 7, 7, 2, 2, 541, 543, 9, 2, 2, 2, 542, 541, 3, 2, 2, 2,
	543, 546, 3, 2, 2, 2, 544, 542, 3, 2, 2, 2, 544, 545, 3, 2, 2, 2, 545,
	555, 3, 2, 2, 2, 546, 544, 3, 2, 2, 2, 547, 549, 5, 56, 29, 2, 548, 550,
	5, 10, 6, 2, 549, 548, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551, 549, 3, 2,
	2, 2, 551, 552, 3, 2, 2, 2, 552, 554, 3, 2, 2, 2, 553, 547, 3, 2, 2, 2,
	554, 557, 3, 2, 2, 2, 555, 553, 3, 2, 2, 2, 555, 556, 3, 2, 2, 2, 556,
	561, 3, 2, 2, 2, 557, 555, 3, 2, 2, 2, 558, 560, 9, 2, 2, 2, 559, 558,
	3, 2, 2, 2, 560, 563, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 561, 562, 3, 2,
	2, 2, 562, 564, 3, 2, 2, 2, 563, 561, 3, 2, 2, 2, 564, 565, 7, 8, 2, 2,
	565, 55, 3, 2, 2, 2, 566, 568, 7, 44, 2, 2, 567, 566, 3, 2, 2, 2, 568,
	571, 3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 569, 570, 3, 2, 2, 2, 570, 572,
	3, 2, 2, 2, 571, 569, 3, 2, 2, 2, 572, 573, 5, 62, 32, 2, 573, 57, 3, 2,
	2, 2, 574, 575, 7, 42, 2, 2, 575, 579, 7, 15, 2, 2, 576, 578, 7, 44, 2,
	2, 577, 576, 3, 2, 2, 2, 578, 581, 3, 2, 2, 2, 579, 577, 3, 2, 2, 2, 579,
	580, 3, 2, 2, 2, 580, 590, 3, 2, 2, 2, 581, 579, 3, 2, 2, 2, 582, 587,
	5, 60, 31, 2, 583, 584, 7, 14, 2, 2, 584, 586, 5, 60, 31, 2, 585, 583,
	3, 2, 2, 2, 586, 589, 3, 2, 2, 2, 587, 585, 3, 2, 2, 2, 587, 588, 3, 2,
	2, 2, 588, 591, 3, 2, 2, 2, 589, 587, 3, 2, 2, 2, 590, 582, 3, 2, 2, 2,
	590, 591, 3, 2, 2, 2, 591, 595, 3, 2, 2, 2, 592, 594, 7, 44, 2, 2, 593,
	592, 3, 2, 2, 2, 594, 597, 3, 2, 2, 2, 595, 593, 3, 2, 2, 2, 595, 596,
	3, 2, 2, 2, 596, 598, 3, 2, 2, 2, 597, 595, 3, 2, 2, 2, 598, 599, 7, 16,
	2, 2, 599, 59, 3, 2, 2, 2, 600, 601, 5, 26, 14, 2, 601, 61, 3, 2, 2, 2,
	602, 603, 7, 42, 2, 2, 603, 607, 7, 15, 2, 2, 604, 606, 7, 44, 2, 2, 605,
	604, 3, 2, 2, 2, 606, 609, 3, 2, 2, 2, 607, 605, 3, 2, 2, 2, 607, 608,
	3, 2, 2, 2, 608, 618, 3, 2, 2, 2, 609, 607, 3, 2, 2, 2, 610, 615, 5, 64,
	33, 2, 611, 612, 7, 14, 2, 2, 612, 614, 5, 64, 33, 2, 613, 611, 3, 2, 2,
	2, 614, 617, 3, 2, 2, 2, 615, 613, 3, 2, 2, 2, 615, 616, 3, 2, 2, 2, 616,
	619, 3, 2, 2, 2, 617, 615, 3, 2, 2, 2, 618, 610, 3, 2, 2, 2, 618, 619,
	3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620, 621, 7, 16, 2, 2, 621, 63, 3, 2,
	2, 2, 622, 624, 7, 44, 2, 2, 623, 622, 3, 2, 2, 2, 624, 627, 3, 2, 2, 2,
	625, 623, 3, 2, 2, 2, 625, 626, 3, 2, 2, 2, 626, 628, 3, 2, 2, 2, 627,
	625, 3, 2, 2, 2, 628, 632, 7, 42, 2, 2, 629, 631, 7, 44, 2, 2, 630, 629,
	3, 2, 2, 2, 631, 634, 3, 2, 2, 2, 632, 630, 3, 2, 2, 2, 632, 633, 3, 2,
	2, 2, 633, 635, 3, 2, 2, 2, 634, 632, 3, 2, 2, 2, 635, 639, 5, 66, 34,
	2, 636, 638, 7, 44, 2, 2, 637, 636, 3, 2, 2, 2, 638, 641, 3, 2, 2, 2, 639,
	637, 3, 2, 2, 2, 639, 640, 3, 2, 2, 2, 640, 656, 3, 2, 2, 2, 641, 639,
	3, 2, 2, 2, 642, 644, 7, 44, 2, 2, 643, 642, 3, 2, 2, 2, 644, 647, 3, 2,
	2, 2, 645, 643, 3, 2, 2, 2, 645, 646, 3, 2, 2, 2, 646, 648, 3, 2, 2, 2,
	647, 645, 3, 2, 2, 2, 648, 652, 5, 66, 34, 2, 649, 651, 7, 44, 2, 2, 650,
	649, 3, 2, 2, 2, 651, 654, 3, 2, 2, 2, 652, 650, 3, 2, 2, 2, 652, 653,
	3, 2, 2, 2, 653, 656, 3, 2, 2, 2, 654, 652, 3, 2, 2, 2, 655, 625, 3, 2,
	2, 2, 655, 645, 3, 2, 2, 2, 656, 65, 3, 2, 2, 2, 657, 664, 7, 42, 2, 2,
	658, 659, 7, 42, 2, 2, 659, 660, 7, 10, 2, 2, 660, 664, 7, 42, 2, 2, 661,
	664, 5, 68, 35, 2, 662, 664, 5, 70, 36, 2, 663, 657, 3, 2, 2, 2, 663, 658,
	3, 2, 2, 2, 663, 661, 3, 2, 2, 2, 663, 662, 3, 2, 2, 2, 664, 67, 3, 2,
	2, 2, 665, 669, 7, 12, 2, 2, 666, 668, 7, 44, 2, 2, 667, 666, 3, 2, 2,
	2, 668, 671, 3, 2, 2, 2, 669, 667, 3, 2, 2, 2, 669, 670, 3, 2, 2, 2, 670,
	675, 3, 2, 2, 2, 671, 669, 3, 2, 2, 2, 672, 676, 7, 42, 2, 2, 673, 676,
	7, 48, 2, 2, 674, 676, 5, 72, 37, 2, 675, 672, 3, 2, 2, 2, 675, 673, 3,
	2, 2, 2, 675, 674, 3, 2, 2, 2, 675, 676, 3, 2, 2, 2, 676, 680, 3, 2, 2,
	2, 677, 679, 7, 44, 2, 2, 678, 677, 3, 2, 2, 2, 679, 682, 3, 2, 2, 2, 680,
	678, 3, 2, 2, 2, 680, 681, 3, 2, 2, 2, 681, 683, 3, 2, 2, 2, 682, 680,
	3, 2, 2, 2, 683, 684, 7, 13, 2, 2, 684, 685, 7, 42, 2, 2, 685, 69, 3, 2,
	2, 2, 686, 687, 7, 21, 2, 2, 687, 688, 5, 66, 34, 2, 688, 689, 7, 13, 2,
	2, 689, 690, 5, 66, 34, 2, 690, 71, 3, 2, 2, 2, 691, 692, 7, 35, 2, 2,
	692, 73, 3, 2, 2, 2, 98, 76, 81, 90, 96, 103, 108, 114, 122, 129, 136,
	141, 146, 154, 161, 165, 171, 178, 186, 190, 196, 202, 209, 217, 222, 230,
	236, 249, 256, 262, 271, 278, 285, 290, 293, 298, 302, 314, 316, 323, 332,
	337, 342, 353, 360, 366, 374, 384, 390, 395, 401, 407, 414, 418, 424, 428,
	434, 440, 447, 454, 459, 461, 466, 470, 475, 482, 488, 494, 502, 509, 516,
	522, 528, 532, 537, 544, 551, 555, 561, 569, 579, 587, 590, 595, 607, 615,
	618, 625, 632, 639, 645, 652, 655, 663, 669, 675, 680,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'module'", "'imports:'", "':'", "'func'", "'{'", "'}'", "'return'",
	"'.'", "'...'", "'['", "']'", "','", "'('", "')'", "'var'", "'type'", "'implements'",
	"'interface'", "'map['", "'||'", "'&&'", "'=='", "'!='", "'<'", "'>'",
	"'<='", "'>='", "'+'", "'-'", "'<<'", "'>>'", "'^'", "'*'", "'/'", "'%'",
	"'='", "'|'", "'&'", "'&^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "OR", "AND", "EQ", "NE", "LT", "GT", "LE", "GE", "PLUS", "MINUS",
	"LSHIFT", "RSHIFT", "CARET", "MUL", "DIV", "REM", "ASSIGN", "BITWISE_OR",
	"BITWISE_AND", "AND_NOT", "IDENTIFIER", "STRING_LIT", "WS", "NEWLINE",
	"LITTLE_U_VALUE", "BIG_U_VALUE", "INT_LIT", "FLOAT_LIT", "IMAGINARY_LIT",
	"RUNE_LIT",
}

var ruleNames = []string{
	"prog", "packageClause", "importDecl", "importSpec", "eos", "topLevelDecl",
	"funcDecl", "methodDecl", "block", "statement", "returnExpr", "statementDecl",
	"expression", "primaryExpr", "listItemExpr", "arraySelection", "operand",
	"literal", "basicLit", "operandName", "qualifiedIdent", "varDecl", "declaration",
	"typeDecl", "structDecl", "typeSpec", "interfaceDecl", "methodSpec", "funcCallSpec",
	"funcCallArg", "funcSpec", "funcArg", "typeRule", "arrayType", "mapType",
	"star",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ConcertoParser struct {
	*antlr.BaseParser
}

func NewConcertoParser(input antlr.TokenStream) *ConcertoParser {
	this := new(ConcertoParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Concerto.G4"

	return this
}

// ConcertoParser tokens.
const (
	ConcertoParserEOF            = antlr.TokenEOF
	ConcertoParserT__0           = 1
	ConcertoParserT__1           = 2
	ConcertoParserT__2           = 3
	ConcertoParserT__3           = 4
	ConcertoParserT__4           = 5
	ConcertoParserT__5           = 6
	ConcertoParserT__6           = 7
	ConcertoParserT__7           = 8
	ConcertoParserT__8           = 9
	ConcertoParserT__9           = 10
	ConcertoParserT__10          = 11
	ConcertoParserT__11          = 12
	ConcertoParserT__12          = 13
	ConcertoParserT__13          = 14
	ConcertoParserT__14          = 15
	ConcertoParserT__15          = 16
	ConcertoParserT__16          = 17
	ConcertoParserT__17          = 18
	ConcertoParserT__18          = 19
	ConcertoParserOR             = 20
	ConcertoParserAND            = 21
	ConcertoParserEQ             = 22
	ConcertoParserNE             = 23
	ConcertoParserLT             = 24
	ConcertoParserGT             = 25
	ConcertoParserLE             = 26
	ConcertoParserGE             = 27
	ConcertoParserPLUS           = 28
	ConcertoParserMINUS          = 29
	ConcertoParserLSHIFT         = 30
	ConcertoParserRSHIFT         = 31
	ConcertoParserCARET          = 32
	ConcertoParserMUL            = 33
	ConcertoParserDIV            = 34
	ConcertoParserREM            = 35
	ConcertoParserASSIGN         = 36
	ConcertoParserBITWISE_OR     = 37
	ConcertoParserBITWISE_AND    = 38
	ConcertoParserAND_NOT        = 39
	ConcertoParserIDENTIFIER     = 40
	ConcertoParserSTRING_LIT     = 41
	ConcertoParserWS             = 42
	ConcertoParserNEWLINE        = 43
	ConcertoParserLITTLE_U_VALUE = 44
	ConcertoParserBIG_U_VALUE    = 45
	ConcertoParserINT_LIT        = 46
	ConcertoParserFLOAT_LIT      = 47
	ConcertoParserIMAGINARY_LIT  = 48
	ConcertoParserRUNE_LIT       = 49
)

// ConcertoParser rules.
const (
	ConcertoParserRULE_prog           = 0
	ConcertoParserRULE_packageClause  = 1
	ConcertoParserRULE_importDecl     = 2
	ConcertoParserRULE_importSpec     = 3
	ConcertoParserRULE_eos            = 4
	ConcertoParserRULE_topLevelDecl   = 5
	ConcertoParserRULE_funcDecl       = 6
	ConcertoParserRULE_methodDecl     = 7
	ConcertoParserRULE_block          = 8
	ConcertoParserRULE_statement      = 9
	ConcertoParserRULE_returnExpr     = 10
	ConcertoParserRULE_statementDecl  = 11
	ConcertoParserRULE_expression     = 12
	ConcertoParserRULE_primaryExpr    = 13
	ConcertoParserRULE_listItemExpr   = 14
	ConcertoParserRULE_arraySelection = 15
	ConcertoParserRULE_operand        = 16
	ConcertoParserRULE_literal        = 17
	ConcertoParserRULE_basicLit       = 18
	ConcertoParserRULE_operandName    = 19
	ConcertoParserRULE_qualifiedIdent = 20
	ConcertoParserRULE_varDecl        = 21
	ConcertoParserRULE_declaration    = 22
	ConcertoParserRULE_typeDecl       = 23
	ConcertoParserRULE_structDecl     = 24
	ConcertoParserRULE_typeSpec       = 25
	ConcertoParserRULE_interfaceDecl  = 26
	ConcertoParserRULE_methodSpec     = 27
	ConcertoParserRULE_funcCallSpec   = 28
	ConcertoParserRULE_funcCallArg    = 29
	ConcertoParserRULE_funcSpec       = 30
	ConcertoParserRULE_funcArg        = 31
	ConcertoParserRULE_typeRule       = 32
	ConcertoParserRULE_arrayType      = 33
	ConcertoParserRULE_mapType        = 34
	ConcertoParserRULE_star           = 35
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) PackageClause() IPackageClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageClauseContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConcertoParserEOF, 0)
}

func (s *ProgContext) ImportDecl() IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *ProgContext) AllTopLevelDecl() []ITopLevelDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem())
	var tst = make([]ITopLevelDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDeclContext)
		}
	}

	return tst
}

func (s *ProgContext) TopLevelDecl(i int) ITopLevelDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDeclContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConcertoParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.PackageClause()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__1 {
		{
			p.SetState(73)
			p.ImportDecl()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConcertoParserT__3)|(1<<ConcertoParserT__15)|(1<<ConcertoParserT__17))) != 0 {
		{
			p.SetState(76)
			p.TopLevelDecl()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(ConcertoParserEOF)
	}

	return localctx
}

// IPackageClauseContext is an interface to support dynamic dispatch.
type IPackageClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageClauseContext differentiates from other interfaces.
	IsPackageClauseContext()
}

type PackageClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageClauseContext() *PackageClauseContext {
	var p = new(PackageClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_packageClause
	return p
}

func (*PackageClauseContext) IsPackageClauseContext() {}

func NewPackageClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageClauseContext {
	var p = new(PackageClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_packageClause

	return p
}

func (s *PackageClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *PackageClauseContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *PackageClauseContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *PackageClauseContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *PackageClauseContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *PackageClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterPackageClause(s)
	}
}

func (s *PackageClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitPackageClause(s)
	}
}

func (s *PackageClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitPackageClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) PackageClause() (localctx IPackageClauseContext) {
	localctx = NewPackageClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConcertoParserRULE_packageClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(ConcertoParserT__0)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(85)
			p.Match(ConcertoParserWS)
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(91)
			p.Eos()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportSpecContext)(nil)).Elem())
	var tst = make([]IImportSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportSpecContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConcertoParserRULE_importDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(ConcertoParserT__1)
	}
	{
		p.SetState(97)
		p.Eos()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(98)
				p.ImportSpec()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(103)
			p.Eos()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *ImportSpecContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ImportSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ImportSpecContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserSTRING_LIT, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConcertoParserRULE_importSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(109)
				p.Match(ConcertoParserWS)
			}

			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(115)
			p.Match(ConcertoParserIDENTIFIER)
		}
		{
			p.SetState(116)
			p.Eos()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(117)
				p.Match(ConcertoParserWS)
			}

			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(123)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(124)
				p.Match(ConcertoParserWS)
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(130)
			p.Match(ConcertoParserT__2)
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(131)
				p.Match(ConcertoParserWS)
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(137)
			p.Match(ConcertoParserSTRING_LIT)
		}
		{
			p.SetState(138)
			p.Eos()
		}

	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, 0)
}

func (s *EosContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *EosContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConcertoParserRULE_eos)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(141)
			p.Match(ConcertoParserWS)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(ConcertoParserNEWLINE)
	}

	return localctx
}

// ITopLevelDeclContext is an interface to support dynamic dispatch.
type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}

type TopLevelDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_topLevelDecl
	return p
}

func (*TopLevelDeclContext) IsTopLevelDeclContext() {}

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_topLevelDecl

	return p
}

func (s *TopLevelDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDeclContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TopLevelDeclContext) FuncDecl() IFuncDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopLevelDeclContext) MethodDecl() IMethodDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodDeclContext)
}

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTopLevelDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TopLevelDecl() (localctx ITopLevelDeclContext) {
	localctx = NewTopLevelDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConcertoParserRULE_topLevelDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.FuncDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.MethodDecl()
		}

	}

	return localctx
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcDecl
	return p
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) MethodSpec() IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *FuncDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConcertoParserRULE_funcDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(ConcertoParserT__3)
	}
	{
		p.SetState(155)
		p.MethodSpec()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(156)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(162)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	{
		p.SetState(165)
		p.Block()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(166)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodDeclContext is an interface to support dynamic dispatch.
type IMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodDeclContext differentiates from other interfaces.
	IsMethodDeclContext()
}

type MethodDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDeclContext() *MethodDeclContext {
	var p = new(MethodDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_methodDecl
	return p
}

func (*MethodDeclContext) IsMethodDeclContext() {}

func NewMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclContext {
	var p = new(MethodDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_methodDecl

	return p
}

func (s *MethodDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *MethodDeclContext) MethodSpec() IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *MethodDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MethodDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *MethodDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *MethodDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *MethodDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *MethodDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterMethodDecl(s)
	}
}

func (s *MethodDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitMethodDecl(s)
	}
}

func (s *MethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitMethodDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) MethodDecl() (localctx IMethodDeclContext) {
	localctx = NewMethodDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConcertoParserRULE_methodDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(ConcertoParserT__3)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(173)
			p.Match(ConcertoParserWS)
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(179)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(180)
		p.MethodSpec()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(181)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(187)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	{
		p.SetState(190)
		p.Block()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(191)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *BlockContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *BlockContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConcertoParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(197)
			p.Match(ConcertoParserWS)
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(ConcertoParserT__4)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(204)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConcertoParserT__6)|(1<<ConcertoParserT__9)|(1<<ConcertoParserT__12)|(1<<ConcertoParserT__14))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ConcertoParserIDENTIFIER-40))|(1<<(ConcertoParserSTRING_LIT-40))|(1<<(ConcertoParserINT_LIT-40))|(1<<(ConcertoParserFLOAT_LIT-40))|(1<<(ConcertoParserIMAGINARY_LIT-40))|(1<<(ConcertoParserRUNE_LIT-40)))) != 0) {
		{
			p.SetState(210)
			p.Statement()
		}
		{
			p.SetState(211)
			p.Eos()
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
			{
				p.SetState(212)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(223)
		p.Match(ConcertoParserT__5)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StatementDecl() IStatementDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementDeclContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) ReturnExpr() IReturnExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnExprContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ConcertoParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.StatementDecl()
		}

	case ConcertoParserT__9, ConcertoParserT__12, ConcertoParserIDENTIFIER, ConcertoParserSTRING_LIT, ConcertoParserINT_LIT, ConcertoParserFLOAT_LIT, ConcertoParserIMAGINARY_LIT, ConcertoParserRUNE_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.expression(0)
		}

	case ConcertoParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
			p.ReturnExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReturnExprContext is an interface to support dynamic dispatch.
type IReturnExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnExprContext differentiates from other interfaces.
	IsReturnExprContext()
}

type ReturnExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnExprContext() *ReturnExprContext {
	var p = new(ReturnExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_returnExpr
	return p
}

func (*ReturnExprContext) IsReturnExprContext() {}

func NewReturnExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnExprContext {
	var p = new(ReturnExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_returnExpr

	return p
}

func (s *ReturnExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ReturnExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ReturnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterReturnExpr(s)
	}
}

func (s *ReturnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitReturnExpr(s)
	}
}

func (s *ReturnExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitReturnExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ReturnExpr() (localctx IReturnExprContext) {
	localctx = NewReturnExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ConcertoParserRULE_returnExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(ConcertoParserT__6)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(231)
			p.Match(ConcertoParserWS)
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(236)
		p.expression(0)
	}

	return localctx
}

// IStatementDeclContext is an interface to support dynamic dispatch.
type IStatementDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementDeclContext differentiates from other interfaces.
	IsStatementDeclContext()
}

type StatementDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementDeclContext() *StatementDeclContext {
	var p = new(StatementDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_statementDecl
	return p
}

func (*StatementDeclContext) IsStatementDeclContext() {}

func NewStatementDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementDeclContext {
	var p = new(StatementDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_statementDecl

	return p
}

func (s *StatementDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementDeclContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStatementDecl(s)
	}
}

func (s *StatementDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStatementDecl(s)
	}
}

func (s *StatementDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStatementDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) StatementDecl() (localctx IStatementDeclContext) {
	localctx = NewStatementDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ConcertoParserRULE_statementDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.VarDecl()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(ConcertoParserOR, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(ConcertoParserAND, 0)
}

func (s *ExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(ConcertoParserEQ, 0)
}

func (s *ExpressionContext) NE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNE, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserLT, 0)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserLE, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserGT, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserGE, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ConcertoParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ConcertoParserMINUS, 0)
}

func (s *ExpressionContext) BITWISE_OR() antlr.TerminalNode {
	return s.GetToken(ConcertoParserBITWISE_OR, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(ConcertoParserCARET, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(ConcertoParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(ConcertoParserDIV, 0)
}

func (s *ExpressionContext) REM() antlr.TerminalNode {
	return s.GetToken(ConcertoParserREM, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserRSHIFT, 0)
}

func (s *ExpressionContext) BITWISE_AND() antlr.TerminalNode {
	return s.GetToken(ConcertoParserBITWISE_AND, 0)
}

func (s *ExpressionContext) AND_NOT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserAND_NOT, 0)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ConcertoParserASSIGN, 0)
}

func (s *ExpressionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ExpressionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ConcertoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ConcertoParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.primaryExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_expression)
			p.SetState(243)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ConcertoParserWS {
				{
					p.SetState(244)
					p.Match(ConcertoParserWS)
				}

				p.SetState(249)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(250)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ConcertoParserOR-20))|(1<<(ConcertoParserAND-20))|(1<<(ConcertoParserEQ-20))|(1<<(ConcertoParserNE-20))|(1<<(ConcertoParserLT-20))|(1<<(ConcertoParserGT-20))|(1<<(ConcertoParserLE-20))|(1<<(ConcertoParserGE-20))|(1<<(ConcertoParserPLUS-20))|(1<<(ConcertoParserMINUS-20))|(1<<(ConcertoParserLSHIFT-20))|(1<<(ConcertoParserRSHIFT-20))|(1<<(ConcertoParserCARET-20))|(1<<(ConcertoParserMUL-20))|(1<<(ConcertoParserDIV-20))|(1<<(ConcertoParserREM-20))|(1<<(ConcertoParserASSIGN-20))|(1<<(ConcertoParserBITWISE_OR-20))|(1<<(ConcertoParserBITWISE_AND-20))|(1<<(ConcertoParserAND_NOT-20)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ConcertoParserWS {
				{
					p.SetState(251)
					p.Match(ConcertoParserWS)
				}

				p.SetState(256)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(257)
				p.expression(2)
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *PrimaryExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) AllListItemExpr() []IListItemExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListItemExprContext)(nil)).Elem())
	var tst = make([]IListItemExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListItemExprContext)
		}
	}

	return tst
}

func (s *PrimaryExprContext) ListItemExpr(i int) IListItemExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListItemExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListItemExprContext)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) FuncCallSpec() IFuncCallSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallSpecContext)
}

func (s *PrimaryExprContext) ArraySelection() IArraySelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraySelectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraySelectionContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *ConcertoParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ConcertoParserRULE_primaryExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__12, ConcertoParserIDENTIFIER, ConcertoParserSTRING_LIT, ConcertoParserINT_LIT, ConcertoParserFLOAT_LIT, ConcertoParserIMAGINARY_LIT, ConcertoParserRUNE_LIT:
		{
			p.SetState(264)
			p.Operand()
		}

	case ConcertoParserT__9:
		{
			p.SetState(265)
			p.Match(ConcertoParserT__9)
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(266)
					p.Match(ConcertoParserWS)
				}

			}
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ConcertoParserT__9 || _la == ConcertoParserT__12 || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ConcertoParserIDENTIFIER-40))|(1<<(ConcertoParserSTRING_LIT-40))|(1<<(ConcertoParserINT_LIT-40))|(1<<(ConcertoParserFLOAT_LIT-40))|(1<<(ConcertoParserIMAGINARY_LIT-40))|(1<<(ConcertoParserRUNE_LIT-40)))) != 0) {
			{
				p.SetState(272)
				p.primaryExpr(0)
			}
			p.SetState(288)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					p.SetState(276)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == ConcertoParserWS {
						{
							p.SetState(273)
							p.Match(ConcertoParserWS)
						}

						p.SetState(278)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}
					{
						p.SetState(279)
						p.ListItemExpr()
					}
					p.SetState(283)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

					for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
						if _alt == 1 {
							{
								p.SetState(280)
								p.Match(ConcertoParserWS)
							}

						}
						p.SetState(285)
						p.GetErrorHandler().Sync(p)
						_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
					}

				}
				p.SetState(290)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
			}

		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(293)
				p.Match(ConcertoParserWS)
			}

			p.SetState(298)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(299)
			p.Match(ConcertoParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(303)
					p.Match(ConcertoParserT__7)
				}
				{
					p.SetState(304)
					p.Match(ConcertoParserIDENTIFIER)
				}

			case 2:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(306)
					p.Match(ConcertoParserT__7)
				}
				{
					p.SetState(307)
					p.FuncCallSpec()
				}

			case 3:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(309)
					p.Match(ConcertoParserT__8)
				}

			case 4:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ConcertoParserRULE_primaryExpr)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(311)
					p.ArraySelection()
				}

			}

		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IListItemExprContext is an interface to support dynamic dispatch.
type IListItemExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListItemExprContext differentiates from other interfaces.
	IsListItemExprContext()
}

type ListItemExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListItemExprContext() *ListItemExprContext {
	var p = new(ListItemExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_listItemExpr
	return p
}

func (*ListItemExprContext) IsListItemExprContext() {}

func NewListItemExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListItemExprContext {
	var p = new(ListItemExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_listItemExpr

	return p
}

func (s *ListItemExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ListItemExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ListItemExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ListItemExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ListItemExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListItemExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListItemExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterListItemExpr(s)
	}
}

func (s *ListItemExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitListItemExpr(s)
	}
}

func (s *ListItemExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitListItemExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ListItemExpr() (localctx IListItemExprContext) {
	localctx = NewListItemExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ConcertoParserRULE_listItemExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(ConcertoParserT__11)
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(318)
			p.Match(ConcertoParserWS)
		}

		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(324)
		p.primaryExpr(0)
	}

	return localctx
}

// IArraySelectionContext is an interface to support dynamic dispatch.
type IArraySelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArraySelectionContext differentiates from other interfaces.
	IsArraySelectionContext()
}

type ArraySelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArraySelectionContext() *ArraySelectionContext {
	var p = new(ArraySelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_arraySelection
	return p
}

func (*ArraySelectionContext) IsArraySelectionContext() {}

func NewArraySelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraySelectionContext {
	var p = new(ArraySelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_arraySelection

	return p
}

func (s *ArraySelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraySelectionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ArraySelectionContext) Star() IStarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *ArraySelectionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ArraySelectionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ArraySelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraySelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterArraySelection(s)
	}
}

func (s *ArraySelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitArraySelection(s)
	}
}

func (s *ArraySelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitArraySelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ArraySelection() (localctx IArraySelectionContext) {
	localctx = NewArraySelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ConcertoParserRULE_arraySelection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(ConcertoParserT__9)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(327)
			p.Match(ConcertoParserWS)
		}

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__9, ConcertoParserT__12, ConcertoParserIDENTIFIER, ConcertoParserSTRING_LIT, ConcertoParserINT_LIT, ConcertoParserFLOAT_LIT, ConcertoParserIMAGINARY_LIT, ConcertoParserRUNE_LIT:
		{
			p.SetState(333)
			p.primaryExpr(0)
		}

	case ConcertoParserMUL:
		{
			p.SetState(334)
			p.Star()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(337)
			p.Match(ConcertoParserWS)
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(343)
		p.Match(ConcertoParserT__10)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) FuncCallSpec() IFuncCallSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallSpecContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *OperandContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *OperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ConcertoParserRULE_operand)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.FuncCallSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.OperandName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(347)
			p.Match(ConcertoParserT__12)
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(348)
				p.Match(ConcertoParserWS)
			}

			p.SetState(353)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(354)
			p.expression(0)
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(355)
				p.Match(ConcertoParserWS)
			}

			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(361)
			p.Match(ConcertoParserT__13)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(363)
			p.Literal()
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ConcertoParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.BasicLit()
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserINT_LIT, 0)
}

func (s *BasicLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserFLOAT_LIT, 0)
}

func (s *BasicLitContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIMAGINARY_LIT, 0)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserRUNE_LIT, 0)
}

func (s *BasicLitContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserSTRING_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterBasicLit(s)
	}
}

func (s *BasicLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitBasicLit(s)
	}
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ConcertoParserRULE_basicLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(ConcertoParserSTRING_LIT-41))|(1<<(ConcertoParserINT_LIT-41))|(1<<(ConcertoParserFLOAT_LIT-41))|(1<<(ConcertoParserIMAGINARY_LIT-41))|(1<<(ConcertoParserRUNE_LIT-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *OperandNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ConcertoParserRULE_operandName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(370)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(371)
			p.QualifiedIdent()
		}

	}

	return localctx
}

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_qualifiedIdent
	return p
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitQualifiedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) QualifiedIdent() (localctx IQualifiedIdentContext) {
	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ConcertoParserRULE_qualifiedIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(375)
		p.Match(ConcertoParserT__7)
	}
	{
		p.SetState(376)
		p.Match(ConcertoParserIDENTIFIER)
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *VarDeclContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *VarDeclContext) FuncCallSpec() IFuncCallSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallSpecContext)
}

func (s *VarDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *VarDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *VarDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ConcertoParserRULE_varDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Match(ConcertoParserT__14)
		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ConcertoParserWS {
			{
				p.SetState(379)
				p.Match(ConcertoParserWS)
			}

			p.SetState(382)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(384)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(385)
				p.Match(ConcertoParserWS)
			}

			p.SetState(390)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(391)
				p.Match(ConcertoParserIDENTIFIER)
			}

		case 2:
			{
				p.SetState(392)
				p.FuncCallSpec()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(ConcertoParserT__14)
		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ConcertoParserWS {
			{
				p.SetState(396)
				p.Match(ConcertoParserWS)
			}

			p.SetState(399)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(401)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(402)
				p.Match(ConcertoParserWS)
			}

			p.SetState(407)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(408)
			p.Match(ConcertoParserASSIGN)
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(409)
				p.Match(ConcertoParserWS)
			}

			p.SetState(414)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(415)
			p.expression(0)
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclarationContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *DeclarationContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ConcertoParserRULE_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.TypeDecl()
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(419)
			p.Eos()
		}

		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) StructDecl() IStructDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *TypeDeclContext) InterfaceDecl() IInterfaceDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ConcertoParserRULE_typeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(426)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)
			p.StructDecl()
		}

	case ConcertoParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(425)
			p.InterfaceDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_structDecl
	return p
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *StructDeclContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *StructDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *StructDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *StructDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *StructDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *StructDeclContext) AllTypeSpec() []ITypeSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem())
	var tst = make([]ITypeSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecContext)
		}
	}

	return tst
}

func (s *StructDeclContext) TypeSpec(i int) ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ConcertoParserRULE_structDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(428)
		p.Match(ConcertoParserT__15)
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(429)
			p.Match(ConcertoParserWS)
		}

		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(434)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(435)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__16 {
		{
			p.SetState(441)
			p.Match(ConcertoParserT__16)
		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(445)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == ConcertoParserWS {
					{
						p.SetState(442)
						p.Match(ConcertoParserWS)
					}

					p.SetState(447)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(448)
					p.Match(ConcertoParserIDENTIFIER)
				}
				p.SetState(452)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())

				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(449)
							p.Match(ConcertoParserWS)
						}

					}
					p.SetState(454)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(457)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
		}

	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(461)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())
	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(467)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(470)
			p.Match(ConcertoParserWS)
		}

		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(476)
		p.Match(ConcertoParserT__4)
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(477)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(483)
				p.TypeSpec()
			}

		}
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext())
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(489)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(495)
		p.Match(ConcertoParserT__5)
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *TypeSpecContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *TypeSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *TypeSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *TypeSpecContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *TypeSpecContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ConcertoParserRULE_typeSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(497)
			p.Match(ConcertoParserWS)
		}

		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(503)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(504)
			p.Match(ConcertoParserWS)
		}

		p.SetState(509)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(510)
		p.TypeRule()
	}
	p.SetState(512)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(511)
				p.Eos()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext())
	}

	return localctx
}

// IInterfaceDeclContext is an interface to support dynamic dispatch.
type IInterfaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceDeclContext differentiates from other interfaces.
	IsInterfaceDeclContext()
}

type InterfaceDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceDeclContext() *InterfaceDeclContext {
	var p = new(InterfaceDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_interfaceDecl
	return p
}

func (*InterfaceDeclContext) IsInterfaceDeclContext() {}

func NewInterfaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclContext {
	var p = new(InterfaceDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_interfaceDecl

	return p
}

func (s *InterfaceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *InterfaceDeclContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *InterfaceDeclContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *InterfaceDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserNEWLINE)
}

func (s *InterfaceDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, i)
}

func (s *InterfaceDeclContext) AllMethodSpec() []IMethodSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem())
	var tst = make([]IMethodSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodSpecContext)
		}
	}

	return tst
}

func (s *InterfaceDeclContext) MethodSpec(i int) IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *InterfaceDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *InterfaceDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *InterfaceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterInterfaceDecl(s)
	}
}

func (s *InterfaceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitInterfaceDecl(s)
	}
}

func (s *InterfaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitInterfaceDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) InterfaceDecl() (localctx IInterfaceDeclContext) {
	localctx = NewInterfaceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ConcertoParserRULE_interfaceDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(516)
		p.Match(ConcertoParserT__17)
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserWS {
		{
			p.SetState(517)
			p.Match(ConcertoParserWS)
		}

		p.SetState(520)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(522)
		p.Match(ConcertoParserIDENTIFIER)
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(523)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())
	}
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserNEWLINE {
		{
			p.SetState(529)
			p.Match(ConcertoParserNEWLINE)
		}

	}
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(532)
			p.Match(ConcertoParserWS)
		}

		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(538)
		p.Match(ConcertoParserT__4)
	}
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(539)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(544)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext())
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(545)
				p.MethodSpec()
			}
			p.SetState(547)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(546)
						p.Eos()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(549)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 75, p.GetParserRuleContext())
			}

		}
		p.SetState(555)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())
	}
	p.SetState(559)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS || _la == ConcertoParserNEWLINE {
		{
			p.SetState(556)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ConcertoParserWS || _la == ConcertoParserNEWLINE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(561)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(562)
		p.Match(ConcertoParserT__5)
	}

	return localctx
}

// IMethodSpecContext is an interface to support dynamic dispatch.
type IMethodSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodSpecContext differentiates from other interfaces.
	IsMethodSpecContext()
}

type MethodSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodSpecContext() *MethodSpecContext {
	var p = new(MethodSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_methodSpec
	return p
}

func (*MethodSpecContext) IsMethodSpecContext() {}

func NewMethodSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodSpecContext {
	var p = new(MethodSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_methodSpec

	return p
}

func (s *MethodSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodSpecContext) FuncSpec() IFuncSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSpecContext)
}

func (s *MethodSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *MethodSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *MethodSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterMethodSpec(s)
	}
}

func (s *MethodSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitMethodSpec(s)
	}
}

func (s *MethodSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitMethodSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) MethodSpec() (localctx IMethodSpecContext) {
	localctx = NewMethodSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ConcertoParserRULE_methodSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(564)
			p.Match(ConcertoParserWS)
		}

		p.SetState(569)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(570)
		p.FuncSpec()
	}

	return localctx
}

// IFuncCallSpecContext is an interface to support dynamic dispatch.
type IFuncCallSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallSpecContext differentiates from other interfaces.
	IsFuncCallSpecContext()
}

type FuncCallSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallSpecContext() *FuncCallSpecContext {
	var p = new(FuncCallSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcCallSpec
	return p
}

func (*FuncCallSpecContext) IsFuncCallSpecContext() {}

func NewFuncCallSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallSpecContext {
	var p = new(FuncCallSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcCallSpec

	return p
}

func (s *FuncCallSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *FuncCallSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncCallSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncCallSpecContext) AllFuncCallArg() []IFuncCallArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncCallArgContext)(nil)).Elem())
	var tst = make([]IFuncCallArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncCallArgContext)
		}
	}

	return tst
}

func (s *FuncCallSpecContext) FuncCallArg(i int) IFuncCallArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncCallArgContext)
}

func (s *FuncCallSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncCallSpec(s)
	}
}

func (s *FuncCallSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncCallSpec(s)
	}
}

func (s *FuncCallSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncCallSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncCallSpec() (localctx IFuncCallSpecContext) {
	localctx = NewFuncCallSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ConcertoParserRULE_funcCallSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(572)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(573)
		p.Match(ConcertoParserT__12)
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(574)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext())
	}
	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__9 || _la == ConcertoParserT__12 || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ConcertoParserIDENTIFIER-40))|(1<<(ConcertoParserSTRING_LIT-40))|(1<<(ConcertoParserINT_LIT-40))|(1<<(ConcertoParserFLOAT_LIT-40))|(1<<(ConcertoParserIMAGINARY_LIT-40))|(1<<(ConcertoParserRUNE_LIT-40)))) != 0) {
		{
			p.SetState(580)
			p.FuncCallArg()
		}
		p.SetState(585)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserT__11 {
			{
				p.SetState(581)
				p.Match(ConcertoParserT__11)
			}
			{
				p.SetState(582)
				p.FuncCallArg()
			}

			p.SetState(587)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(590)
			p.Match(ConcertoParserWS)
		}

		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(596)
		p.Match(ConcertoParserT__13)
	}

	return localctx
}

// IFuncCallArgContext is an interface to support dynamic dispatch.
type IFuncCallArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallArgContext differentiates from other interfaces.
	IsFuncCallArgContext()
}

type FuncCallArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallArgContext() *FuncCallArgContext {
	var p = new(FuncCallArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcCallArg
	return p
}

func (*FuncCallArgContext) IsFuncCallArgContext() {}

func NewFuncCallArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallArgContext {
	var p = new(FuncCallArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcCallArg

	return p
}

func (s *FuncCallArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FuncCallArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncCallArg(s)
	}
}

func (s *FuncCallArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncCallArg(s)
	}
}

func (s *FuncCallArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncCallArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncCallArg() (localctx IFuncCallArgContext) {
	localctx = NewFuncCallArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ConcertoParserRULE_funcCallArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(598)
		p.expression(0)
	}

	return localctx
}

// IFuncSpecContext is an interface to support dynamic dispatch.
type IFuncSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSpecContext differentiates from other interfaces.
	IsFuncSpecContext()
}

type FuncSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSpecContext() *FuncSpecContext {
	var p = new(FuncSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcSpec
	return p
}

func (*FuncSpecContext) IsFuncSpecContext() {}

func NewFuncSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSpecContext {
	var p = new(FuncSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcSpec

	return p
}

func (s *FuncSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *FuncSpecContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncSpecContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncSpecContext) AllFuncArg() []IFuncArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncArgContext)(nil)).Elem())
	var tst = make([]IFuncArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncArgContext)
		}
	}

	return tst
}

func (s *FuncSpecContext) FuncArg(i int) IFuncArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncArgContext)
}

func (s *FuncSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncSpec(s)
	}
}

func (s *FuncSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncSpec(s)
	}
}

func (s *FuncSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncSpec() (localctx IFuncSpecContext) {
	localctx = NewFuncSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ConcertoParserRULE_funcSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(600)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(601)
		p.Match(ConcertoParserT__12)
	}
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(602)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(607)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())
	}
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConcertoParserT__9 || _la == ConcertoParserT__18 || _la == ConcertoParserIDENTIFIER || _la == ConcertoParserWS {
		{
			p.SetState(608)
			p.FuncArg()
		}
		p.SetState(613)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserT__11 {
			{
				p.SetState(609)
				p.Match(ConcertoParserT__11)
			}
			{
				p.SetState(610)
				p.FuncArg()
			}

			p.SetState(615)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(618)
		p.Match(ConcertoParserT__13)
	}

	return localctx
}

// IFuncArgContext is an interface to support dynamic dispatch.
type IFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgContext differentiates from other interfaces.
	IsFuncArgContext()
}

type FuncArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgContext() *FuncArgContext {
	var p = new(FuncArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_funcArg
	return p
}

func (*FuncArgContext) IsFuncArgContext() {}

func NewFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgContext {
	var p = new(FuncArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_funcArg

	return p
}

func (s *FuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *FuncArgContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *FuncArgContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *FuncArgContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) FuncArg() (localctx IFuncArgContext) {
	localctx = NewFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ConcertoParserRULE_funcArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(653)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 91, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(623)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(620)
				p.Match(ConcertoParserWS)
			}

			p.SetState(625)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(626)
			p.Match(ConcertoParserIDENTIFIER)
		}
		p.SetState(630)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(627)
				p.Match(ConcertoParserWS)
			}

			p.SetState(632)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(633)
			p.TypeRule()
		}
		p.SetState(637)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(634)
				p.Match(ConcertoParserWS)
			}

			p.SetState(639)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(643)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(640)
				p.Match(ConcertoParserWS)
			}

			p.SetState(645)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(646)
			p.TypeRule()
		}
		p.SetState(650)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConcertoParserWS {
			{
				p.SetState(647)
				p.Match(ConcertoParserWS)
			}

			p.SetState(652)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITypeRuleContext is an interface to support dynamic dispatch.
type ITypeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRuleContext differentiates from other interfaces.
	IsTypeRuleContext()
}

type TypeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRuleContext() *TypeRuleContext {
	var p = new(TypeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *TypeRuleContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *TypeRuleContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *TypeRuleContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (s *TypeRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitTypeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) TypeRule() (localctx ITypeRuleContext) {
	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ConcertoParserRULE_typeRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(661)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 92, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(655)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(656)
			p.Match(ConcertoParserIDENTIFIER)
		}
		{
			p.SetState(657)
			p.Match(ConcertoParserT__7)
		}
		{
			p.SetState(658)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(659)
			p.ArrayType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(660)
			p.MapType()
		}

	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *ArrayTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *ArrayTypeContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserWS)
}

func (s *ArrayTypeContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserWS, i)
}

func (s *ArrayTypeContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserINT_LIT, 0)
}

func (s *ArrayTypeContext) Star() IStarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ConcertoParserRULE_arrayType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(663)
		p.Match(ConcertoParserT__9)
	}
	p.SetState(667)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(664)
				p.Match(ConcertoParserWS)
			}

		}
		p.SetState(669)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext())
	}
	p.SetState(673)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConcertoParserIDENTIFIER:
		{
			p.SetState(670)
			p.Match(ConcertoParserIDENTIFIER)
		}

	case ConcertoParserINT_LIT:
		{
			p.SetState(671)
			p.Match(ConcertoParserINT_LIT)
		}

	case ConcertoParserMUL:
		{
			p.SetState(672)
			p.Star()
		}

	case ConcertoParserT__10, ConcertoParserWS:

	default:
	}
	p.SetState(678)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserWS {
		{
			p.SetState(675)
			p.Match(ConcertoParserWS)
		}

		p.SetState(680)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(681)
		p.Match(ConcertoParserT__10)
	}
	{
		p.SetState(682)
		p.Match(ConcertoParserIDENTIFIER)
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) AllTypeRule() []ITypeRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem())
	var tst = make([]ITypeRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeRuleContext)
		}
	}

	return tst
}

func (s *MapTypeContext) TypeRule(i int) ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ConcertoParserRULE_mapType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(684)
		p.Match(ConcertoParserT__18)
	}
	{
		p.SetState(685)
		p.TypeRule()
	}
	{
		p.SetState(686)
		p.Match(ConcertoParserT__10)
	}
	{
		p.SetState(687)
		p.TypeRule()
	}

	return localctx
}

// IStarContext is an interface to support dynamic dispatch.
type IStarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStarContext differentiates from other interfaces.
	IsStarContext()
}

type StarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarContext() *StarContext {
	var p = new(StarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_star
	return p
}

func (*StarContext) IsStarContext() {}

func NewStarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarContext {
	var p = new(StarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_star

	return p
}

func (s *StarContext) GetParser() antlr.Parser { return s.parser }
func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStar(s)
	}
}

func (s *StarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStar(s)
	}
}

func (s *StarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Star() (localctx IStarContext) {
	localctx = NewStarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ConcertoParserRULE_star)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(689)
		p.Match(ConcertoParserMUL)
	}

	return localctx
}

func (p *ConcertoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 13:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ConcertoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ConcertoParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
