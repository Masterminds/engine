{{ define "globalAttrs" }}{{ with .Id}}id="{{.}}"
{{end}}{{with .AccessKey}}accesskey="{{.}}"
{{end}}{{with .Dir}}dir="{{.}}"
{{end}}{{with .Lang}}lang="{{.}}"
{{end}}{{with .Role}}role="{{.}}"
{{end}}{{with .Style}}style="{{.}}"
{{end}}{{with .TabIndex}}tabindex="{{.}}"
{{end}}{{with .Title}}title="{{.}}"
{{end}}{{with .Translate}}translate="{{.}}"
{{end}}{{if eq 1 .ContentEditable}}contenteditable="true"{{else if eq 2 .ContentEditable }}contenteditable="false"
{{end}}{{if .Hidden | eq 1}}hidden="true"{{else if .Hidden | eq 2 }}hidden="false"
{{end}}{{with .Class}}class="{{join " " .}}"
{{end}}{{with .Aria}}{{range $k, $v := .}}{{$k}}="{{$v}}"{{end}}
{{end}}{{with .Data}}{{range $k, $v := .}}{{$k}}="{{$v}}"{{end}}{{end}}{{end}}

{{define "form.button"}}<button {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Value}}value="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .Menu}}menu="{{.}}"
{{end}}{{with .Type}}type="{{.}}"
{{end}}{{if .Autofocus}}autofocus="true"
{{end}}{{if .Disabled}}disabled="true"
{{end}}>{{end}}

{{define "form.keygen"}}<keygen {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{if .Autofocus}}autofocus="true"
{{end}}{{if .Disabled}}disabled="true"
{{end}}{{with .KeyType}}keytype="{{.}}"
{{end}}{{with .Challenge}}challenge="{{.}}"{{end}}>{{end}}

{{define "form.label"}}<label {{template "globalAttrs" .}}{{with .For}}for="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}>{{.Text}}</label>
{{end}}

{{define "form.output"}}<output {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .For}}for="{{.}}"{{end}}>{{end}}


{{define "form.progress"}}<progress {{template "globalAttrs" .}}{{with .Value}}value="{{.}}"
{{end}}{{with .Max}}max="{{.}}"{{end}}>{{end}}

{{define "form.meter"}}<meter {{template "globalAttrs" .}}{{with .Value}}value="{{.}}"
{{end}}{{with .Min}}min="{{.}}"
{{end}}{{with .Max}}max="{{.}}"
{{end}}{{with .Low}}low="{{.}}"
{{end}}{{with .High}}high="{{.}}"
{{end}}{{with .Optimum}}optimum="{{.}}"{{end}}>{{end}}

{{define "form.option"}}<option {{template "globalAttrs" .}}{{with .Selected}}selected
{{end}}{{with .Value}}value="{{.}}"
{{end}}{{with .Disabled}}disabled{{end}}>{{.Label | default .Value}}</option>
{{end}}

{{define "form.optgroup"}}<optgroup {{template "globalAttrs" .}}{{with .Disabled}}disabled
{{end}}{{with .Label}}label="{{.}}"{{end}}>
{{range .Options}}{{template "form.option" .}}{{end}}</optgroup>{{end}}

{{define "form.optitems"}}
{{if . | typeIsLike "form.Option" }}{{template "form.option" . }}{{end}}
{{if . | typeIsLike "form.OptGroup" }}{{template "form.optgroup" . }}{{end}}
{{end}}

{{define "form.datalist"}}<datalist {{template "globalAttrs" .}}>
{{range .Options}}{{template "form.option" .}}{{end}}</datalist>{{end}}

{{define "form.select"}}
{{if len .Label | lt 0}}<label for="{{.Name}}">{{.Label}}</label>{{end}}
<select {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Autofocus}}autofocus
{{end}}{{with .Disabled}}disabled
{{end}}{{with .Multiple}}multiple
{{end}}{{with .Required}}required
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{with .Size}}size="{{.}}"{{end}}>{{range .Options}}
{{template "form.optitems" .}}
{{end}}</select>{{end}}

{{define "form.textarea"}}<textarea {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
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
{{define "form.text"}}{{template "form.input" .}}{{end}}
{{define "form.password"}}{{template "form.input" .}}{{end}}
{{define "form.submit"}}{{template "form.input" .}}{{end}}
{{define "form.tel"}}{{template "form.input" .}}{{end}}
{{define "form.url"}}{{template "form.input" .}}{{end}}
{{define "form.email"}}{{template "form.input" .}}{{end}}
{{define "form.date"}}{{template "form.input" .}}{{end}}
{{define "form.time"}}{{template "form.input" .}}{{end}}
{{define "form.number"}}{{template "form.input" .}}{{end}}
{{define "form.range"}}{{template "form.input" .}}{{end}}
{{define "form.color"}}{{template "form.input" .}}{{end}}
{{define "form.checkbox"}}{{template "form.input" .}}{{end}}
{{define "form.radio"}}{{template "form.input" .}}{{end}}
{{define "form.file"}}{{template "form.input" .}}{{end}}
{{define "form.image"}}{{template "form.input" .}}{{end}}
{{define "form.reset"}}{{template "form.input" .}}{{end}}
{{define "form.hidden"}}{{template "form.input" .}}{{end}}

{{define "form.buttoninput"}}
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

{{define "form.input"}}
{{if len .Label | lt 0}}<label for="{{.Name}}">.Label</label>
{{end}}<input type="{{$t := typeOf . | split "."}}{{lower $t._1}}" {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
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



{{define "form.fieldloop"}}
{{range .}}
<p>{{typeOf .}}</p>
{{if . | typeIsLike "form.Button" }}{{template "form.button" . }}{{end}}
{{if . | typeIsLike "form.FieldSet" }}{{template "form.fieldset" . }}{{end}}
{{if . | typeIsLike "form.Label" }}{{template "form.label" . }}{{end}}
{{if . | typeIsLike "form.Keygen" }}{{template "form.keygen" . }}{{end}}
{{if . | typeIsLike "form.Output" }}{{template "form.output" . }}{{end}}
{{if . | typeIsLike "form.Progress" }}{{template "form.progress" . }}{{end}}
{{if . | typeIsLike "form.Meter" }}{{template "form.meter" . }}{{end}}
{{if . | typeIsLike "form.DataList" }}{{template "form.datalist" . }}{{end}}
{{if . | typeIsLike "form.Select" }}{{template "form.select" . }}{{end}}
{{if . | typeIsLike "form.TextArea" }}{{template "form.textarea" . }}{{end}}
{{if . | typeIsLike "form.Input" }}{{template "form.input" . }}{{end}}
{{if . | typeIsLike "form.Password" }}{{template "form.password" . }}{{end}}
{{if . | typeIsLike "form.Text" }}{{template "form.text" . }}{{end}}
{{if . | typeIsLike "form.Submit" }}{{template "form.submit" . }}{{end}}
{{if . | typeIsLike "form.Tel" }}{{template "form.tel" . }}{{end}}
{{if . | typeIsLike "form.URL" }}{{template "form.url" . }}{{end}}
{{if . | typeIsLike "form.Email" }}{{template "form.email" . }}{{end}}
{{if . | typeIsLike "form.Date" }}{{template "form.date" . }}{{end}}
{{if . | typeIsLike "form.Time" }}{{template "form.time" . }}{{end}}
{{if . | typeIsLike "form.Number" }}{{template "form.number" . }}{{end}}
{{if . | typeIsLike "form.Range" }}{{template "form.range" . }}{{end}}
{{if . | typeIsLike "form.Color" }}{{template "form.color" . }}{{end}}
{{if . | typeIsLike "form.Checkbox" }}{{template "form.checkbox" . }}{{end}}
{{if . | typeIsLike "form.Radio" }}{{template "form.radio" . }}{{end}}
{{if . | typeIsLike "form.File" }}{{template "form.file" . }}{{end}}
{{if . | typeIsLike "form.Image" }}{{template "form.image" . }}{{end}}
{{if . | typeIsLike "form.Reset" }}{{template "form.range" . }}{{end}}
{{if . | typeIsLike "form.ButtonInput" }}{{template "form.buttoninput" . }}{{end}}
{{if . | typeIsLike "form.Hidden" }}{{template "form.hidden" . }}{{end}}
{{if . | typeIsLike "form.Div" }}{{template "form.div" .}}{{end}}
{{end}}
{{end}}

{{define "form.fieldset"}}
<fieldset {{template "globalAttrs" .}}{{with .Name}}name="{{.}}"
{{end}}{{with .Form}}form="{{.}}"
{{end}}{{if .Disabled}}disabled="true"
{{end}}>{{with .Legend}}<legend>{{.}}</legend>
{{end}}{{template "form.fieldloop" .Fields}}
</fieldset>{{end}}

{{define "form.div"}}<div {{template "globalAttrs" .}}>{{template "form.fieldloop" .Fields}}</div>{{end}}


{{define "form"}}
<form {{template "globalAttrs" .  }}{{with .Name}}name="{{.}}" {{end}}
{{with .AcceptCharset}}acceptchars="{{.}}"
{{end}}{{with .Enctype}}enctype="{{.}}"
{{end}}{{with .Action }}action="{{.}}"
{{end}}{{with .Method}}method="{{.}}"
{{end}}{{with .Target}}target="{{.}}"
{{end}}{{with .Autocomplete}}autocomplete="true"
{{end}}{{with .Novalidate}}novalidate="true" {{end}}>
{{template "form.fieldloop" .Fields}}
</form>
{{end}}

