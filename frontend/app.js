async function verifyEmail() {
    const email = document.getElementById("emailInput").value;
    const resultDiv = document.getElementById("result");

    if (!email) {
        resultDiv.innerHTML = "❌ Please enter an email";
        return;
    }

    resultDiv.innerHTML = "⏳ Verifying...";

    try {
        const response = await fetch(`http://localhost:8080/verify?email=${email}`);
        const data = await response.json();

        resultDiv.innerHTML = `
            <p><strong>Email:</strong> ${data.email}</p>
            <p>✅ Syntax: ${data.valid_syntax}</p>
            <p>🌐 Domain: ${data.valid_domain}</p>
            <p>📡 SMTP: ${data.smtp_valid}</p>
            <p>🚫 Disposable: ${data.is_disposable}</p>
            <p>📬 MX: ${data.mx_records}</p>
        `;
    } catch (error) {
        resultDiv.innerHTML = "❌ Error connecting to server";
    }
}