{{ define "globalAttrs" }}{{ with .Id}}id="{{.}}"
{{end}}{{with .AccessKey}}accesskey="{{.}}"
{{end}}{{with .Dir}}dir="{{.}}"
{{end}}{{with .Lang}}lang="{{.}}"
{{end}}{{with .Style}}style="{{.}}"
{{end}}{{with .TabIndex}}tabindex="{{.}}"
{{end}}{{with .Title}}title="{{.}}"
{{end}}{{with .Translate}}translate="{{.}}"
{{end}}{{if .ContentEditable | eq 1}}contenteditable="true"{{else if .ContentEditable | eq 2 }}contenteditable="false"
{{end}}{{if .Hidden | eq 1}}hidden="true"{{else if .Hidden | eq 2 }}hidden="false"
{{end}}{{with .Class}}class="{{join " " .}}"
{{end}}{{with .Data}}{{range $k, $v := .}}{{$k}}="{{$v}}"{{end}}{{end}}{{end}}

<form {{template "globalAttrs" .  }}{{with .Name}}name="{{.}}" {{end}}
{{with .AcceptCharset}}acceptchars="{{.}}"
{{end}}{{with .Enctype}}enctype="{{.}}"
{{end}}{{with .Action }}action="{{.}}"
{{end}}{{with .Method}}method="{{.}}"
{{end}}{{with .Target}}target="{{.}}"
{{end}}{{with .Autocomplete}}autocomplete="true"
{{end}}{{with .Novalidate}}novalidate="true" {{end}}>
{{range .Fields}}
<p>{{.Name}}</p>
{{end}}
</form>

