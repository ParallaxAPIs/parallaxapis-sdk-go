# 🚀 Parallax SDK: Datadome & Perimeterx (Go)

Easily interact with Datadome and Perimeterx anti-bot solutions using a simple Go SDK. Fast integration, clear API! ✨

---

## 📦 Installation

```bash
go get github.com/parallaxsystems/parallax-sdk-go
```

---

## 🧑‍💻 Datadome Usage

### ⚡ SDK Initialization

```go
import "github.com/parallaxsystems/parallax-sdk-go"

// Basic initialization with API key
sdk := parallaxsdk.NewDatadomeSDK("key", "")

// Custom host
sdk := parallaxsdk.NewDatadomeSDK("key", "example.host.com")
```

### 🕵️‍♂️ Generate New User Agent

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
userAgent, err := sdk.GenerateUserAgent(parallaxsdk.TaskGenUserAgent{
    Region: "pl",
    Site: "vinted",
})
if err != nil {
    panic(err)
}
fmt.Println(userAgent)
// Output:
// &parallaxsdk.UserAgentResponse{
//     Message: "New device successfully created.",
//     UserAgent: "Mozilla/5.0 ...",
//     SecHeader: "...",
//     SecFullVersionList: "...",
//     SecPlatform: "...",
//     SecArch: "...",
// }
```

### 🔍 Get Task Data

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
challengeURL := "https://geo.captcha-delivery.com/captcha/?initialCid=initialCid&cid=cid&referer=referer&hash=hash&t=t&s=s&e=e"
cookie := "cookie"
taskData, productType, err := parallaxsdk.ParseChallengeURL(challengeURL, cookie)
if err != nil {
    panic(err)
}
fmt.Println(taskData, productType)
// Output:
// &parallaxsdk.TaskDatadomeCookieData{
//     Cid: "cookie",
//     B: "",
//     E: "e",
//     S: "s",
//     InitialCid: "initialCid",
// }, "captcha"
```

### 🍪 Generate Cookie

```go
sdk := parallaxsdk.NewDatadomeSDK("key", "")
challengeURL := "https://geo.captcha-delivery.com/captcha/?initialCid=initialCid&cid=cid&referer=referer&hash=hash&t=t&s=s&e=e"
cookie := "cookie"
taskData, productType, err := parallaxsdk.ParseChallengeURL(challengeURL, cookie)
if err != nil {
    panic(err)
}
cookieResp, err := sdk.GenerateDatadomeCookie(parallaxsdk.TaskDatadomeCookie{
    Site: "vinted",
    Region: "pl",
    Data: *taskData,
    Pd: productType,
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
})
if err != nil {
    panic(err)
}
fmt.Println(cookieResp)
// Output:
// &parallaxsdk.DatadomeCookieResponse{
//     Message: "datadome=cookie_value",
//     UserAgent: "Mozilla/5.0 ...",
// }
```

---

## 🛡️ Perimeterx Usage

### ⚡ SDK Initialization

```go
import "github.com/parallaxsystems/parallax-sdk-go"

// Basic initialization with API key
sdk := parallaxsdk.NewPerimeterxSDK("key", "")

// Custom host
sdk := parallaxsdk.NewPerimeterxSDK("key", "example.host.com")
```

### 🍪 Generate PX Cookie

```go
sdk := parallaxsdk.NewPerimeterxSDK("key", "")
result, err := sdk.GenerateCookies(parallaxsdk.TaskGeneratePXCookies{
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Region: "com",
    Site: "stockx",
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", result)
// Output:
// &parallaxsdk.PxCookieResponse{
//     Cookie: "_px3=d3sswjaltwxgAd...",
//     Vid: "514d7e11-6962-11f0-810f-88cc16043287",
//     Cts: "514d8e28-6962-11f0-810f-51b6xf2786b0",
//     IsFlagged: false,
//     IsMaybeFlagged: true,
//     UserAgent: "Mozilla/5.0 ...",
//     Data: "==WlrBti6vpO6rshP1CFtBsiocoO8...",
// }

holdCaptchaResult, err := sdk.GenerateHoldCaptcha(parallaxsdk.TaskGenerateHoldCaptcha{
    Proxy: "http://user:pas@addr:port",
    Proxyregion: "eu",
    Region: "com",
    Site: "stockx",
    Data: result.Data,
    PowPro: "",
})
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", holdCaptchaResult)
// Output:
// &parallaxsdk.PxCookieResponse{
//     Cookie: "_px3=d3sswjaltwxgAd...",
//     Vid: "514d7e11-6962-11f0-810f-88cc16043287",
//     FlaggedPow: false,
// }
```

---

## 📚 Documentation & Help

- Full API docs: [GitHub](https://github.com/parallaxsystems/parallax-sdk-go)
- Issues & support: [GitHub Issues](https://github.com/parallaxsystems/parallax-sdk-go/issues)

---

## 📝 License

MIT

---

Made with ❤️ by Parallax Systems
