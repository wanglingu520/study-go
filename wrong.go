package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Song   string
	Length time.Duration
	Style  string
}

var tracks = []Track{
	{"流行", "周杰伦", "七里香", length("3m38s"), ".mp3"},
	{"流行", "周杰伦", "晴天", length("3m37s"), ".wav"},
	{"抒情", "王菲", "梦中人", length("4m36s"), ".mp3"},
	{"说唱", "潘玮柏", "快乐崇拜", length("4m24s"), ".wav"},
	{"流行", "周杰伦", "七里香", length("3m38s"), ".mp3"},
	{"流行", "周杰伦", "晴天", length("3m37s"), ".wav"},
	{"抒情", "王菲", "梦中人", length("4m36s"), ".mp3"},
	{"说唱", "潘玮柏", "快乐崇拜", length("4m24s"), ".wav"},
}

func length(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

func printTracks(tracks []Track) string {
	const format = "%v\t%v\t%v\t%v\t%v\t<br/>"
	var trackString string
	trackString = fmt.Sprintf(format, "Title", "Artist", "Song", "time", "Style")
	trackString = fmt.Sprintf("%s%s", trackString, fmt.Sprintf(format, "_____", "_____", "____", "____", "_____"))
	for _, t := range tracks {
		trackString = fmt.Sprintf("%s%s", trackString, fmt.Sprintf(format, t.Title, t.Artist, t.Song, t.Length, t.Style))
	}
	return trackString
}

type byArtist []Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//var arrayInt []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pn := r.FormValue("pn")
	pageNumber, err := strconv.Atoi(pn)
	if err != nil {
		return
	}

	arr, errStr := paging((pageNumber-1)*3+1, 3, tracks)
	if errStr != "" && errStr != "已经到了最后一页" {
		return
	}

	data := printTracks(arr)
	str := `<html>
<title>This for Love</title>
<body style="background: #FFB5B5">
<header>
<h1>
<iframe width="420" scrolling="no" height="60" frameborder="0" allowtransparency="true" src="http://i.tianqi.com/index.php?c=code&id=12&icon=1&num=5&site=12"></iframe>
<br>
<b>Beauliful girl</b> 
<b>I love it</b><br>


</h1>
<abbr title="wanglingu520">W.</abbr>
</header>
<h2>
鸟山明
</h2>
<img src="img/wukong.jpg">
<p>
他的成名之作《阿拉蕾》和代表作《龙珠》在全世界都极受欢迎，广受赞誉。其作品在集英社的漫画杂志《周刊少年JUMP》
上连载，常年人气顺位排名第一，绝对的王牌主力之一，十分受各个年龄阶层的读者喜欢，也影响了众多后辈的漫画家、导演等人物。
</p>
<p>
其中《龙珠》更是风靡全球，创造了远超过2亿4000万册的销售量（2016年）。
1984年开始连载的《龙珠》和当时许多连载时红遍天下，完结后却迅速退热的作品不同，
</p>
<pre>
《龙珠》的大热在1995年漫画完结后仍一直延续下去，至今依然有着极高人气。
<pre>

%s<br>

<form>
username:<input><br>
password:<input><br>
<button>signup</button>
</form>
<footer>
<hr>
<a href="http://192.168.1.2:8080/?pn=1" target="_self">1</a>|<a href="http://192.168.1.2:8080/?pn=2" target="_blank">2</a>|<a href="http://192.168.1.2:8080/?pn=3" target="_self">3</a>
</footer>
</body>
</html>`
	w.Write([]byte(fmt.Sprintf(str, data)))

}

func main() {

	http.HandleFunc("/", IndexHandler)

	err := http.ListenAndServe("192.168.1.2:8080", nil)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func paging(page_element_start int, page_size int, array []Track) ([]Track, string) {
	//	var arrayErr []int
	var arraySlice []Track
	pageEnd := page_size + page_element_start - 1
	switch {
	case page_element_start > len(array):
		return arraySlice, "初始元素位置大于查询数组长度"
	case page_element_start+page_size-1 > len(array):
		arraySlice = array[page_element_start-1 : len(array)]
		return arraySlice, "已经到了最后一页"
	//	case page_size > len(array):
	//	arraySlice = array
	case page_element_start < len(array) && page_element_start+page_size-1 <= len(array) && page_size <= len(array):
		arraySlice = array[page_element_start-1 : pageEnd]
		return arraySlice, ""
	}
	return arraySlice, "未知情况"
}

func paging_map_interface(page_element_start int, page_size int, array []map[string]interface{}) ([]map[string]interface{}, string) {
	//	var arrayErr []int
	arraySlice := []map[string]interface{}{}
	pageEnd := page_size + page_element_start - 1
	switch {
	case page_element_start > len(array):
		return arraySlice, "初始元素位置大于查询数组长度"
	case page_element_start+page_size-1 > len(array):
		arraySlice = array[page_element_start-1 : len(array)]
		return arraySlice, "已经到了最后一页"
	//	case page_size > len(array):
	//	arraySlice = array
	case page_element_start < len(array) && page_element_start+page_size-1 <= len(array) && page_size <= len(array):
		arraySlice = array[page_element_start-1 : pageEnd]
		return arraySlice, ""
	}
	return arraySlice, "未知情况"
}
