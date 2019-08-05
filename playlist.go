/*

Your music player allows you to choose songs to play, but only in pairs and only pairs of songs with durations that add up to whole minutes. Given a list of song durations, calculate the total number of different pairs of songs that can be chosen.



For example, given song lengths [40, 20, 60], one pair of songs can be chosen: [40, 20]. While the third song is a whole minute long, songs must be chosen in pairs.



Function Description

Complete the function playlist in the editor below. The function must return a single integer that is the number of ways to choose two songs such that the total duration is a multiple of a whole minute.



playlist has the following parameter(s):

    songs[songs[0],...songs[n-1]]:  an array of integers wherein each element denotes duration of a song in seconds



Constraints

    1 ≤ n ≤ 105
    1 ≤ songs[i] ≤ 1000, where 0 ≤ i < n


Input Format For Custom Testing
Sample Case 0

Sample Input For Custom Testing

4
10
50
90
30

Sample Output

2

Explanation

The first and second songs pair to 60 seconds.  The third and fourth songs pair to 120 seconds.  No other pairs will satisfy the requirement.
Sample Case 1

Sample Input For Custom Testing

5
30
20
150
100
40

Sample Output

3

Explanation

We can select three pairs of songs whose whole duration is a multiple of a whole minute. They are (1, 3), (2, 4) and (2, 5).
Sample Case 2

Sample Input For Custom Testing

3
60
60
60

Sample Output

3

Explanation

There are three pairs of songs that end in whole minutes. They are (1, 2), (1, 3) and (2, 3).
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'playlist' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY songs as parameter.
 */

func playlist(songs []int32) int64 {
	// Write your code here
	hm := map[int32]int{}
	uniq := []int{}
	for _, n := range songs { //collect unique members and count frequency
		if _, ok := hm[n]; !ok {
			uniq = append(uniq, int(n))
		}
		hm[n]++
	}
	result := [][]int{}
	for i := 0; i < len(uniq); i++ {
		for j := i + 1; j < len(uniq); j++ {
			if (uniq[i]+uniq[j])%60 == 0 {
				result = append(result, []int{uniq[i], uniq[j]})
			}
		}
	}
	countDup := 0
	for k, v := range hm {
		if v > 1 {
			if (k+k)%60 == 0 {
				countDup += v * (v - 1) / 2
			}
		}
	}
	count := 0
	for _, p := range result {
		count += hm[int32(p[0])] * hm[int32(p[1])]
	}
	return int64(count + countDup)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	songsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var songs []int32

	for i := 0; i < int(songsCount); i++ {
		songsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		songsItem := int32(songsItemTemp)
		songs = append(songs, songsItem)
	}

	result := playlist(songs)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
