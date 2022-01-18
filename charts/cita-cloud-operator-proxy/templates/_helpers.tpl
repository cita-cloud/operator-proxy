{{/*
Expand the name of the chart.
*/}}
{{- define "cita-cloud-operator-proxy.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "cita-cloud-operator-proxy.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "cita-cloud-operator-proxy.labels" -}}
helm.sh/chart: {{ include "cita-cloud-operator-proxy.chart" . }}
{{ include "cita-cloud-operator-proxy.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "cita-cloud-operator-proxy.selectorLabels" -}}
app.kubernetes.io/name: {{ include "cita-cloud-operator-proxy.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
