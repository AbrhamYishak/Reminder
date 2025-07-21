chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message === "LOGIN_WITH_GOOGLE") {
    chrome.identity.getAuthToken({ interactive: true }, (token) => {
      if (chrome.runtime.lastError) {
        sendResponse({ error: chrome.runtime.lastError.message });
      } else {
        sendResponse({ token });
      }
    });
    return true;
  }
});
