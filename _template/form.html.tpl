{{ define "globalAttrs" }}{{ with .Id}}id="{{.}}"
{{end}}{{with .AccessKey}}accesskey="{{.}}"
{{end}}{{with .Dir}}dir="{{.}}"
{{end}}{{with .Lang}}lang="{{.}}"
{{end}}{{with .Style}}style="{{.}}"
{{end}}{{with .TabIndex}}tabindex="{{.}}"
{{end}}{{with .Title}}title="{{.}}"
{{end}}{{with .Translate}}translate="{{.}}"
{{end}}{{if eq 1 .ContentEditable}}contenteditable="true"{{else if eq 2 .ContentEditable }}contenteditable="false"
{{end}}{{if .Hidden | eq 1}}hidden="true"{{else if .Hidden | eq 2 }}hidden="false"
{{end}}{{with .Class}}class="{{join " " .}}"
{{end}}{{with .Data}}{{range $k, $v := .}}{{$k}}="{{$v}}"{{end}}{{end}}{{end}}

{{define "button"}}<button {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Value}}value="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .Menu}}menu="{{.}}"
{{end}}{{with .Type}}type="{{.}}"
{{end}}{{if .Autofocus}}autofocus="true"
{{end}}{{if .Disabled}}disabled="true"
{{end}}>{{end}}

{{define "keygen"}}<keygen {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{if .Autofocus}}autofocus="true"
{{end}}{{if .Disabled}}disabled="true"
{{end}}{{with .KeyType}}keytype="{{.}}"
{{end}}{{with .Challenge}}challenge="{{.}}"{{end}}>{{end}}

{{define "label"}}<label {{template "globalAttrs" .}}{{with .For}}for="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}>{{.Text}}</label>
{{end}}

{{define "output"}}<output {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .For}}for="{{.}}"{{end}}>{{end}}


{{define "progress"}}<progress {{template "globalAttrs" .}}{{with .Value}}value="{{.}}"
{{end}}{{with .Max}}max="{{.}}"{{end}}>{{end}}

{{define "meter"}}<meter {{template "globalAttrs" .}}{{with .Value}}value="{{.}}"
{{end}}{{with .Min}}min="{{.}}"
{{end}}{{with .Max}}max="{{.}}"
{{end}}{{with .Low}}low="{{.}}"
{{end}}{{with .High}}high="{{.}}"
{{end}}{{with .Optimum}}optimum="{{.}}"{{end}}>{{end}}

{{define "option"}}<option {{template "globalAttrs" .}}{{with .Selected}}selected
{{end}}{{with .Value}}value="{{.}}"
{{end}}{{with .Disabled}}disabled{{end}}>{{.Label | default .Value}}</option>
{{end}}

{{define "optgroup"}}<optgroup {{template "globalAttrs" .}}{{with .Disabled}}disabled
{{end}}{{with .Label}}label="{{.}}"{{end}}>
{{range .Options}}{{template "option" .}}{{end}}</optgroup>{{end}}

{{define "optitems"}}
{{if . | typeIsLike "form.Option" }}{{template "option" . }}{{end}}
{{if . | typeIsLike "form.OptGroup" }}{{template "optgroup" . }}{{end}}
{{end}}

{{define "datalist"}}<datalist {{template "globalAttrs" .}}>
{{range .Options}}{{template "option" .}}{{end}}</datalist>{{end}}

{{define "select"}}
{{if len .Label | lt 0}}<label for="{{.Name}}">{{.Label}}</label>{{end}}
<select {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Autofocus}}autofocus
{{end}}{{with .Disabled}}disabled
{{end}}{{with .Multiple}}multiple
{{end}}{{with .Required}}required
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .Size}}size="{{.}}"{{end}}>{{range .Options}}
{{template "optitems" .}}
{{end}}</select>{{end}}

{{define "textarea"}}<textarea {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Autocomplete}}autocomplete="{{.}}"
{{end}}{{with .Dirname}}dirname="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .Placeholder}}placeholder="{{.}}"
{{end}}{{with .Wrap}}wrap="{{.}}"
{{end}}{{with .Cols}}cols="{{.}}"
{{end}}{{with .MaxLength}}maxlength="{{.}}"
{{end}}{{with .MinLength}}minlength="{{.}}"
{{end}}{{with .Rows}}rows="{{.}}"
{{end}}{{with .Autofocus}}autofocus
{{end}}{{with .Disabled}}disabled
{{end}}{{with .ReadOnly}}readonly
{{end}}{{with .Required}}required{{end}}>{{.Value}}</textarea>{{end}}


{{define "fieldloop"}}
{{range .}}
<p>{{typeOf .}}</p>
{{if . | typeIsLike "form.Button" }}{{template "button" . }}{{end}}
{{if . | typeIsLike "form.FieldSet" }}{{template "fieldset" . }}{{end}}
{{if . | typeIsLike "form.Label" }}{{template "label" . }}{{end}}
{{if . | typeIsLike "form.Keygen" }}{{template "keygen" . }}{{end}}
{{if . | typeIsLike "form.Output" }}{{template "output" . }}{{end}}
{{if . | typeIsLike "form.Progress" }}{{template "progress" . }}{{end}}
{{if . | typeIsLike "form.Meter" }}{{template "meter" . }}{{end}}
{{if . | typeIsLike "form.DataList" }}{{template "datalist" . }}{{end}}
{{if . | typeIsLike "form.Select" }}{{template "select" . }}{{end}}
{{if . | typeIsLike "form.TextArea" }}{{template "textarea" . }}{{end}}
{{end}}
{{end}}

{{define "fieldset"}}
<fieldset {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{if .Disabled}}disabled="true"
{{end}}>{{with .Legend}}<legend>{{.}}</legend>
{{end}}{{template "fieldloop" .Fields}}
{{end}}

<form {{template "globalAttrs" .  }}{{with .Name}}name="{{.}}" {{end}}
{{with .AcceptCharset}}acceptchars="{{.}}"
{{end}}{{with .Enctype}}enctype="{{.}}"
{{end}}{{with .Action }}action="{{.}}"
{{end}}{{with .Method}}method="{{.}}"
{{end}}{{with .Target}}target="{{.}}"
{{end}}{{with .Autocomplete}}autocomplete="true"
{{end}}{{with .Novalidate}}novalidate="true" {{end}}>
{{template "fieldloop" .Fields}}
</form>

