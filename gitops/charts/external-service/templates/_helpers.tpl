{{/* Name of the resources */}}
{{- define "resource-name" -}}
{{ .Release.Name }}
{{- end -}}

{{/* Name of the apps */}}
{{- define "app-name" -}}
{{ default .Release.Name .Values.app }}
{{- end -}}

{{/* Labels of the resources and apps */}}
{{- define "labels" -}}
{{- if not (empty .Values.env) }}
env: {{ .Values.env }}
{{- end }}
app: {{ include "app-name" . }}
{{- end -}}
