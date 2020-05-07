<p>Packages:</p>
<ul>
<li>
<a href="#cilium.networking.extensions.gardener.cloud%2fv1alpha1">cilium.networking.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="cilium.networking.extensions.gardener.cloud/v1alpha1">cilium.networking.extensions.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the configuration of the Cilium Network Extension.</p>
</p>
Resource Types:
<ul><li>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>
</li></ul>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig
</h3>
<p>
<p>NetworkConfig is a struct representing the configmap for the cilium
networking plugin</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
cilium.networking.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>NetworkConfig</code></td>
</tr>
<tr>
<td>
<code>debug</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Debug configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>prometheus</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Prometheus">
Prometheus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Prometheus configuration</p>
</td>
</tr>
<tr>
<td>
<code>operatorprometheus</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.OperatorPrometheus">
OperatorPrometheus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>OperatorPrometheus configuration</p>
</td>
</tr>
<tr>
<td>
<code>psp</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>PSPEnabled configuration</p>
</td>
</tr>
<tr>
<td>
<code>kubeproxy</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxy">
KubeProxy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>KubeProxy configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>hubble</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Hubble">
Hubble
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Hubble configuration to be enabled or not</p>
</td>
</tr>
<tr>
<td>
<code>tunnel</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.TunnelMode">
TunnelMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TunnelMode configuration, it should be &lsquo;vxlan&rsquo;, &lsquo;geneve&rsquo; or &lsquo;disabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>store</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Store">
Store
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Store can be either Kubernetes or etcd.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.ExternalIP">ExternalIP
</h3>
<p>
<p>ExternalIPs configuration for cilium</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>externalipEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>ExternalIPenabled is used to define whether ExternalIP address is required or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Hubble">Hubble
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Hubble enablement for cilium</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ui</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>metrics</code></br>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.IdentityAllocationMode">IdentityAllocationMode
(<code>string</code> alias)</p></h3>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxy">KubeProxy
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>KubeProxy configuration for cilium</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disabled</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enabled specifies whether kubeproxy is disabled.</p>
</td>
</tr>
<tr>
<td>
<code>k8sServiceHost</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceHost specify the controlplane node IP Address.</p>
</td>
</tr>
<tr>
<td>
<code>k8sServicePort</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServicePort specify the kube-apiserver port number.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.KubeProxyReplacementMode">KubeProxyReplacementMode
(<code>string</code> alias)</p></h3>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.NodePortMode">NodePortMode
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.Nodeport">Nodeport</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Nodeport">Nodeport
</h3>
<p>
<p>Nodeport enablement for cilium</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>nodePortEnabled</code></br>
<em>
bool
</em>
</td>
<td>
<p>Enabled is used to define whether Nodeport is required or not.</p>
</td>
</tr>
<tr>
<td>
<code>nodePortMode</code></br>
<em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NodePortMode">
NodePortMode
</a>
</em>
</td>
<td>
<p>Mode is the mode of NodePort feature</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.OperatorPrometheus">OperatorPrometheus
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>OperatorPrometheus configuration for cilium</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>operatorprometheusEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Prometheus">Prometheus
</h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>Prometheus configuration for cilium</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>prometheusEnabled</code></br>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.Store">Store
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
</p>
<h3 id="cilium.networking.extensions.gardener.cloud/v1alpha1.TunnelMode">TunnelMode
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#cilium.networking.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
</p>
<hr/>