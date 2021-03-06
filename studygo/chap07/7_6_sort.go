package main

import (
	"time"
	"os"
	"text/tabwriter"
	"fmt"
)

// 排序的三个要素：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式
type Interfac interface {
	Len()			int
	Less(i, j int)	bool
	Swap(i, j int)
}

type StringSlice []string

func (p StringSlice) Len() int { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i]

}

//下面的变量tracks包好了一个播放列表

type Track struct {
	Title	string
	Artist	string
	Album	string
	Year	int
	Length	time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
    {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
    {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// printTracks函数将播放列表打印成一个表格
func printTracks(tracks []*Track)  {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artists", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()	// calculate column widths and print table
}

// 为了能按照Artist字段对播放列表进行排序，我们会像对StringSlice那样定义一个新的带有必须Len，Less和Swap方法的切片类型。
type byArtist []*Track
func (x byArtist) Len() int { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int) { x[i], x[j] = x[j], x[i] }