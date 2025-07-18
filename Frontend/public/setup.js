document.getElementById('setupForm').addEventListener('submit', async function (e) {
  e.preventDefault();

  const timezone = document.getElementById('timezone').value;
  const token = localStorage.getItem("ReminderToken");

  try {
    const res = await fetch('http://localhost:8080/setup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify({ TimeZone: timezone })
    });

    const data = await res.json();
    if (res.ok) {
      alert('Setup Completed.');

      const popupUrl = chrome.runtime.getURL('index.html');

      window.location.href = popupUrl;

      if (chrome.tabs && chrome.tabs.query) {
        chrome.tabs.query({ active: true, currentWindow: true }, function(tabs) {
          if (tabs[0] && tabs[0].id) {
            chrome.tabs.remove(tabs[0].id);
          }
        });
     }
	 alert("Open the Extension")
    } else {
      alert(data.message || 'Setup failed');
    }
  } catch (err) {
    console.error(err);
    alert('An error occurred. Please check your network.');
  }
});
