# <img src="https://github.com/ParallaxAPIs/.github/blob/main/profile/logo.png" alt="Parallax Logo" width="30" height="30" style="vertical-align: middle;"> ParallaxAPIs Go SDK: Datadome & PerimeterX

**Go SDK for bypassing DataDome and PerimeterX anti-bot protection.**

## 📖 Overview

ParallaxAPIs provides a **request-based solution** for bypassing DataDome and PerimeterX anti-bot systems. Instead of relying on slow, resource-heavy browser automation, our API generates valid cookies and tokens in **200-400ms** through direct HTTP requests.

**What We Solve:**
- ✅ **DataDome** - Slider captchas, interstitial pages, cookie generation, tags payload
- ✅ **PerimeterX** - Cookie generation (_px3), challenge solver, vid & cts tokens

**Key Benefits:**
- ⚡ **Lightning Fast** - 200-400ms response times vs 5-10+ seconds for browsers
- 🔧 **Simple Integration** - Clean API with comprehensive documentation, no browser management required
- 🚀 **Highly Scalable** - Handle thousands of concurrent requests with minimal resources
- ⚙️ **Flexible Configuration** - Custom timeouts, HTTP clients, and proxy settings
- 💰 **Cost Effective** - Lightweight infrastructure, minimal proxy usage
- 🔄 **Always Updated** - We handle all reverse engineering and updates for you

---

## 🚀 Quick Start

Get started with ParallaxAPIs SDK's in under 5 minutes:

1. **Join our [Discord](https://www.parallaxsystems.io/join?s=gh)** - Connect with our community
2. **Create a ticket** - Request your API key
3. **Get your free trial** - Start testing immediately
4. **[Install the SDK](#-installation)** - Choose your preferred language
5. **Solve all anti-bots in seconds** - Start bypassing DataDome, PerimeterX & more

---

## 📦 Installation

```bash
go get github.com/ParallaxAPIs/parallaxapis-sdk-go
```

![Go Get Demo](https://raw.githubusercontent.com/ParallaxAPIs/.github/main/profile/gogetparallax.gif)

---

## 🧑‍💻 Datadome Usage

### ⚡ SDK Initialization

```go
import (
    "time"
    "github.com/ParallaxAPIs/parallaxapis-sdk-go"
)

// Basic initialization with API key
sdk := parallaxsdk.NewDatadomeSDK("Key", "")

// Custom host
sdk := parallaxsdk.NewDatadomeSDK("Key", "https://example.host.com")

// With custom timeout (default is 30 seconds)
sdk := parallaxsdk.NewDatadomeSDK("Key", "", parallaxsdk.WithCustomTimeout(60*time.Second))

// With HTTP proxy for client requests
sdk := parallaxsdk.NewDatadomeSDK("Key", "", parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"))

// Multiple options combined
sdk := parallaxsdk.NewDatadomeSDK("Key", "https://example.host.com",
    parallaxsdk.WithCustomTimeout(45*time.Second),
    parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"),
    parallaxsdk.WithInsecureSkipVerify(),
)

usage, err := sdk.CheckUsage("site")
if err != nil {
    fmt.Println("Error checking usage:", err)
    return
}
fmt.Println(usage)
```

### 🕵️‍♂️ Generate New User Agent

```go
sdk := parallaxsdk.NewDatadomeSDK("Key", "")

userAgent, err := sdk.GenerateUserAgent(parallaxsdk.TaskGenUserAgent{
    Region: "com",
    Site: "site",
})
if err != nil {
    panic(err)
}

fmt.Println(userAgent)
```

### 🔍 Get Task Data

```go
sdk := parallaxsdk.NewDatadomeSDK("Key", "")

challengeURL := "https://www.example.com/captcha/?initialCid=initialCid&cid=cid&referer=referer&hash=hash&t=t&s=1&e=e"
cookie := "cookie_value"

taskData, productType, err := parallaxsdk.ParseChallengeURL(challengeURL, cookie)
if err != nil {
    panic(err)
}

fmt.Println(taskData, productType)
```

### 📄 Parse Challenge HTML

```go
htmlBody := "<html><script>dd={example:1}</script></html>"
prevCookie := "cookie_value"

taskData, productType, err := parallaxsdk.ParseChallengeHTML(htmlBody, prevCookie)
if err != nil {
    panic(err)
}

fmt.Println(taskData, productType)
```

### 🍪 Generate Cookie

```go
sdk := parallaxsdk.NewDatadomeSDK("Key", "")

challengeURL := "https://www.example.com/captcha/?initialCid=initialCid&cid=cid&referer=referer&hash=hash&t=t&s=1&e=e"
cookie := "cookie_value"

taskData, productType, err := parallaxsdk.ParseChallengeURL(challengeURL, cookie)
if err != nil {
    panic(err)
}

cookieResp, err := sdk.GenerateDatadomeCookie(parallaxsdk.TaskDatadomeCookie{
    Site: "site",
    Region: "com",
    Data: *taskData,
    Pd: productType,
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
})
if err != nil {
    panic(err)
}

fmt.Println(cookieResp)
```

### 🏷️ Generate Tags Cookie

```go
sdk := parallaxsdk.NewDatadomeSDK("Key", "")

cookieResp, err := sdk.GenerateDatadomeTagsCookie(parallaxsdk.TaskDatadomeTagsCookie{
    Site: "site",
    Region: "com",
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Cid: "cookie_value"
})
if err != nil {
    panic(err)
}

fmt.Println(cookieResp)
```

### 🔍 Detect and Parse Challenge

```go
sdk := parallaxsdk.NewDatadomeSDK("Key", "")

responseBody := "<html>...</html>" // Response body from website
prevCookie := "cookie_value"

isBlocked, taskData, productType, err := parallaxsdk.DetectChallengeAndParse(responseBody, prevCookie)
if err != nil {
    panic(err)
}

if isBlocked {
    cookieResp, err := sdk.GenerateDatadomeCookie(parallaxsdk.TaskDatadomeCookie{
        Site: "site",
        Region: "com",
        Data: *taskData,
        Pd: productType,
        Proxy: "http://user:pas@addr:port",
        Proxyregion: "eu",
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(cookieResp)
}
```

---

## 🛡️ Perimeterx Usage

### ⚡ SDK Initialization

```go
import (
    "time"
    "github.com/ParallaxAPIs/parallaxapis-sdk-go"
)

// Basic initialization with API key
sdk := parallaxsdk.NewPerimeterxSDK("Key", "")

// Custom host
sdk := parallaxsdk.NewPerimeterxSDK("Key", "example.host.com")

// With custom timeout (default is 30 seconds)
sdk := parallaxsdk.NewPerimeterxSDK("Key", "", parallaxsdk.WithCustomTimeout(60*time.Second))

// With HTTP proxy for client requests
sdk := parallaxsdk.NewPerimeterxSDK("Key", "", parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"))

// Multiple options combined
sdk := parallaxsdk.NewPerimeterxSDK("Key", "example.host.com",
    parallaxsdk.WithCustomTimeout(45*time.Second),
    parallaxsdk.WithClientProxy("http://user:pass@proxy.example.com:8080"),
    parallaxsdk.WithInsecureSkipVerify(),
)

usage, err := sdk.CheckUsage("site")
if err != nil {
    fmt.Println("Error checking usage:", err)
    return
}
fmt.Println(usage)
```

### 🍪 Generate PX Cookie

```go
sdk := parallaxsdk.NewPerimeterxSDK("Key", "")

result, err := sdk.GenerateCookies(parallaxsdk.TaskGeneratePXCookies{
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Region: "com",
    Site: "site",
})
if err != nil {
    panic(err)
}

fmt.Printf(result)


holdCaptchaResult, err := sdk.GenerateHoldCaptcha(parallaxsdk.TaskGenerateHoldCaptcha{
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Region: "com",
    Site: "site",
    Data: result.Data,
    PowPro: "",
})
if err != nil {
    panic(err)
}

fmt.Printf(holdCaptchaResult)
```

---

## 📚 Documentation & Help

- Full API docs & support: [Discord](https://www.parallaxsystems.io/join?s=gh)



## 🌟 Contributing

Got feedback or found a bug? Feel free to open an issue or send us a pull request!



## 🏢 Enterprise

Unlock enterprise-grade performance with custom solutions, expanded limits, and expert support. [Contact us](https://www.parallaxsystems.io/join?s=gh) to learn more.



## 📝 License

MIT

---

## 🔑 Keywords

**DataDome bypass** • **PerimeterX bypass** • **Anti-bot bypass** • **Bot detection bypass** • **CAPTCHA solver** • **Cookie generator** • **Go web scraping** • **Go bot automation** • **Golang anti-bot** • **DataDome Go SDK** • **PerimeterX Go SDK** • **Headless browser alternative** • **Request-based bypass** • **Go automation** • **Web scraping Go** • **Bot mitigation bypass** • **Sensor data generation** • **Challenge solver**
