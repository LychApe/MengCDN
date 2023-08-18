::: tip
FallSoft提供了两套CDN服务
:::

## 使用声明
1.我们无法完全保证CDN的SLA。<br>
2.请切勿对CDN进行任何攻击测试，任何未经允许的渗透均为违法行为。<br>
3.我们会尽力保证CDN的SLA，但是我们也无法保证100%的正常运行时间，所以当我们无法向提供加速服务之时，我们会通知您。<br>
4.在主要为中国大陆提供的CDN中，为避免滥用，我们采用白名单模式，您需要向我们提交申请，同时该CDN严禁上传或缓存违反中国大陆法律的违规文件。<br>
5.在主要为中国境外提供的CDN中，我们使用了Cloudflare的CDN,您需要遵守Github、Unpkg、WordPress、Cloudflare的规定上传文件。<br>
6.一切解释权归FallSoft和泊集所有。<br>

## 主要面向中国大陆的CDN

首先，您应该向"my@luox.in"，该邮箱申请白名单仓库，或者向CDNJS提交Pull Request。<br>
::: warning
单文件大小小于或等于8MB<br>
单仓库、软件包大小小于或等于2GB<br>
不得上传违反中国大陆法律的内容<br>
:::

我们对该套CDN进行了比较好的优化，我们在中国大陆地区使用了较好的CDN，在中国境外我们使用了Cloudflare CDN以保障全球访问速度。

您可以通过以下的方式调用你的仓库、软件包、WP主题:


### Github仓库

Proxy From: https://raw.githubusercontent.com

```
https://cdn.fallsoft.cn/gh/<user>/<repo>@<version>/<path>
```

### NPM软件包

Proxy From:
https://unpkg.com

```
https://cdn.fallsoft.cn/npm/<package>/<version>/<path>
```

### WordPress

Theme Proxy From: https://themes.svn.wordpress.org/ <br>
Plugin Proxy From: https://plugins.svn.wordpress.org/

```
# For themes in wordpress.org
https://cdn.fallsoft.cn/wp/theme/<name>/<version>/<path>

# For plugins in wordpress.org
https://cdn.fallsoft.cn/wp/plugin/<name>/<version>/<path>
```

## 主要面向中国境外的CDN
::: tip
该CDN无需申请（同时无大小限制）即可使用，但其在中国大陆的加载速度略慢。
:::
这套CDN是完全公开的，所有人的仓库和软件包都被允许加速，无需申请，其使用Cloudflare路由将文件通过Workers缓存到边缘网络，获得较好的速度。

其用法与境内的相同，但其直接分别Proxy了Github Raw、Unpkg、WordPress。

### Github仓库

Proxy From: https://raw.githubusercontent.com

```
https://cdn.iloli.icu/gh/<user>/<repo>/<version>/<path>
```

### NPM软件包

Proxy From:
https://unpkg.com

```
### Full
https://cdn.iloli.icu/npm/@<user>/<package>@<version>/<path>

### No version
https://cdn.iloli.icu/npm/@<user>/<package>/<path>
# equal to
https://cdn.iloli.icu/npm/@<user>/<package>@latest/<path>

### Standard package
https://cdn.iloli.icu/npm/<package>@<version>/<path>

### No path
https://cdn.iloli.icu/npm/<package>@<version>
# will be redirected to
https://cdn.iloli.icu/npm/@<user>/<package>@<version>/<path> # main file

```

### WordPress

Theme Proxy From: https://themes.svn.wordpress.org/ <br>
Plugin Proxy From: https://plugins.svn.wordpress.org/

```
# For themes in wordpress.org
https://cdn.iloli.icu/wp/theme/<name>/<version>/<path>

# For plugins in wordpress.org
https://cdn.iloli.icu/wp/plugin/<name>/<version>/<path>
```

