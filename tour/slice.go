package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    matrix := make([][]uint8, dx)
    for i := 0; i < dx; i ++ {
        matrix[i] = make([]uint8, dy)
        for j := 0; j < dy; j++ {
            matrix[i][j] = uint8((i + j) / 2)
        }
    }
    return matrix
}

func main() {
    pic.Show(Pic)
}
