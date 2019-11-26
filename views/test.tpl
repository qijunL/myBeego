<html>
    <head>
        <title>
            {{.Title}}
        </title>
    </head>
    <body>
         {{.Title}}
         {{range .Users}}
            {{.Username}}{{$.len}}<br>
         {{end}}
    </body>
</html>