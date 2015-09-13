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

{{/* We define a template for each so that overrides are easy. */}}
{{define "text"}}{{template "input" .}}{{end}}
{{define "password"}}{{template "input" .}}{{end}}
{{define "submit"}}{{template "input" .}}{{end}}
{{define "tel"}}{{template "input" .}}{{end}}
{{define "url"}}{{template "input" .}}{{end}}
{{define "email"}}{{template "input" .}}{{end}}
{{define "date"}}{{template "input" .}}{{end}}
{{define "time"}}{{template "input" .}}{{end}}
{{define "number"}}{{template "input" .}}{{end}}
{{define "range"}}{{template "input" .}}{{end}}
{{define "color"}}{{template "input" .}}{{end}}
{{define "checkbox"}}{{template "input" .}}{{end}}
{{define "radio"}}{{template "input" .}}{{end}}
{{define "file"}}{{template "input" .}}{{end}}
{{define "image"}}{{template "input" .}}{{end}}
{{define "reset"}}{{template "input" .}}{{end}}
{{define "hidden"}}{{template "input" .}}{{end}}

{{define "buttoninput"}}
{{if len .Label | lt 0}}<label for="{{.Name}}">.Label</label>
{{end}}<input type="button" {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Accept}}accept="{{.}}"
{{end}}{{with .Alt}}alt="{{.}}"
{{end}}{{with .Autocomplete}}autocomplete="{{.}}"
{{end}}{{with .Dirname}}dirname="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .InputMode}}inputmode="{{.}}"
{{end}}{{with .Placeholder}}placeholder="{{.}}"
{{end}}{{with .Src}}src="{{.}}"
{{end}}{{with .Value}}value="{{.}}"
{{end}}{{with .Height}}height="{{.}}"
{{end}}{{with .Width}}width="{{.}}"
{{end}}{{with .Size}}size="{{.}}"
{{end}}{{with .Autofocus}}autofocus
{{end}}{{with .Checked}}checked
{{end}}{{with .Disabled}}disabled
{{end}}{{with .Required}}required
{{end}}>{{end}}

{{define "input"}}
{{if len .Label | lt 0}}<label for="{{.Name}}">.Label</label>
{{end}}<input type="{{typeOf . | lower}}" {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Accept}}accept="{{.}}"
{{end}}{{with .Alt}}alt="{{.}}"
{{end}}{{with .Autocomplete}}autocomplete="{{.}}"
{{end}}{{with .Dirname}}dirname="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .List}}list="{{.}}"
{{end}}{{with .InputMode}}inputmode="{{.}}"
{{end}}{{with .Min}}min="{{.}}"
{{end}}{{with .Max}}max="{{.}}"
{{end}}{{with .MaxLength}}maxlength="{{.}}"
{{end}}{{with .Pattern}}pattern="{{.}}"
{{end}}{{with .Placeholder}}placeholder="{{.}}"
{{end}}{{with .Src}}src="{{.}}"
{{end}}{{with .Step}}step="{{.}}"
{{end}}{{with .Value}}value="{{.}}"
{{end}}{{with .Height}}height="{{.}}"
{{end}}{{with .Width}}width="{{.}}"
{{end}}{{with .Size}}size="{{.}}"
{{end}}{{with .Autofocus}}autofocus
{{end}}{{with .Checked}}checked
{{end}}{{with .Disabled}}disabled
{{end}}{{with .Multiple}}multiple
{{end}}{{with .ReadOnly}}readonly
{{end}}{{with .Required}}required
{{end}}>{{end}}


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
{{if . | typeIsLike "form.Input" }}{{template "input" . }}{{end}}
{{if . | typeIsLike "form.Password" }}{{template "password" . }}{{end}}
{{if . | typeIsLike "form.Text" }}{{template "text" . }}{{end}}
{{if . | typeIsLike "form.Submit" }}{{template "submit" . }}{{end}}
{{if . | typeIsLike "form.Tel" }}{{template "tel" . }}{{end}}
{{if . | typeIsLike "form.URL" }}{{template "url" . }}{{end}}
{{if . | typeIsLike "form.Email" }}{{template "email" . }}{{end}}
{{if . | typeIsLike "form.Date" }}{{template "date" . }}{{end}}
{{if . | typeIsLike "form.Time" }}{{template "time" . }}{{end}}
{{if . | typeIsLike "form.Number" }}{{template "number" . }}{{end}}
{{if . | typeIsLike "form.Range" }}{{template "range" . }}{{end}}
{{if . | typeIsLike "form.Color" }}{{template "color" . }}{{end}}
{{if . | typeIsLike "form.Checkbox" }}{{template "checkbox" . }}{{end}}
{{if . | typeIsLike "form.Radio" }}{{template "radio" . }}{{end}}
{{if . | typeIsLike "form.File" }}{{template "file" . }}{{end}}
{{if . | typeIsLike "form.Image" }}{{template "image" . }}{{end}}
{{if . | typeIsLike "form.Reset" }}{{template "range" . }}{{end}}
{{if . | typeIsLike "form.ButtonInput" }}{{template "buttoninput" . }}{{end}}
{{if . | typeIsLike "form.Hidden" }}{{template "hidden" . }}{{end}}
{{end}}
{{end}}

{{define "fieldset"}}
<fieldset {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{if .Disabled}}disabled="true"
{{end}}>{{with .Legend}}<legend>{{.}}</legend>
{{end}}{{template "fieldloop" .Fields}}
{{end}}

{{define "form"}}
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
{{end}}

