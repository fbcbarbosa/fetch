package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Example() {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(pudimWebPage))
		}))
	defer srv.Close()

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", srv.URL}
	main()
	// Output:
	// <html>
	// <html xmlns="http://www.w3.org/1999/xhtml">
	// <head>
	//     <title>Pudim</title>
	//     <link rel="stylesheet" href="estilo.css">
	// </head>
	// <body>
	// <div>
	//     <div class="container">
	//         <div class="image">
	//             <img src="pudim.jpg" alt="">
	//         </div>
	//         <div class="email">
	//             <a href="mailto:pudim@pudim.com.br">pudim@pudim.com.br</a>
	//         </div>
	//     </div>
	// </div>
	// <script>
	//     (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	//                 (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	//             m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	//     })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
	//
	//     ga('create', 'UA-28861757-1', 'auto');
	//     ga('send', 'pageview');
	//
	// </script>
	// </body>
	// </html>
}

func TestFetch_local(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(pudimWebPage))
		}))
	defer srv.Close()

	var tests = []struct {
		url string
	}{
		{srv.URL},
		{strings.Trim(srv.URL, "http://")},
	}
	for _, test := range tests {
		buf := new(bytes.Buffer)
		err := Fetch(test.url, buf)
		if err != nil {
			t.Fatalf("Fetch(%q, buffer):\n%v", test.url, err)
		}
		if buf.String() != pudimWebPage {
			t.Errorf("\n\tExpected: \n%q\n\tGot: %q", pudimWebPage, buf.String())
		}
	}
}

func TestFetch_remote(t *testing.T) {
	if testing.Short() {
		t.Skip("Test ignored [option -test.short]")
	}

	var tests = []struct {
		url      string
		contains string
	}{
		{"http://www.pudim.com.br", pudimWebPage},
		{"https://status.github.com/", "GitHub System Status1"},
	}
	for _, test := range tests {
		buf := new(bytes.Buffer)
		err := Fetch(test.url, buf)
		if err != nil {
			t.Fatalf("Fetch(%q, buffer):\n%v", test.url, err)
		}
		if !strings.Contains(buf.String(), test.contains) {
			t.Errorf("\n\t%q\n\tShould contain %q", buf.String()[0:200]+"...", test.contains)
		}
	}
}

const pudimWebPage = `<html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>Pudim</title>
    <link rel="stylesheet" href="estilo.css">
</head>
<body>
<div>
    <div class="container">
        <div class="image">
            <img src="pudim.jpg" alt="">
        </div>
        <div class="email">
            <a href="mailto:pudim@pudim.com.br">pudim@pudim.com.br</a>
        </div>
    </div>
</div>
<script>
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
                (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
            m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
    })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

    ga('create', 'UA-28861757-1', 'auto');
    ga('send', 'pageview');

</script>
</body>
</html>
`
