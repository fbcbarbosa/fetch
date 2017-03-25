package main

import (
	"net/http"
	"net/http/httptest"
	"os"
)

func Example() {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(pudim))
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

var pudim = `<html>
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
