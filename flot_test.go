package lttb

import (
	"bytes"
	"html/template"
)

var reportTmpl = template.Must(template.New("report").Parse(`
<html>
<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.0.3/jquery.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/flot/0.8.2/jquery.flot.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/flot/0.8.2/jquery.flot.time.min.js"></script>

<script type="text/javascript">

    var data = {{ . }};

    $(document).ready(function() {
        $.plot($("#placeholder"), [data], {
                lines: { show: true },
                points: { show: true, fillColor: false },
            })
    })

</script>

<body>

<div id="placeholder" style="width:900px; height:300px"></div>

</body>
</html>
`))

func flot(data [][2]float64) ([]byte, error) {
	var b bytes.Buffer
	err := reportTmpl.Execute(&b, data)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
