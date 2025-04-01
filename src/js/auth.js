function updateAuthState() {
    const username = localStorage.getItem('username');
    const loginLinks = document.querySelectorAll('.login-btn');
    
    loginLinks.forEach(link => {
        if (username) {
            link.textContent = username;
            link.href = '#';
            // Add logout option
            link.onclick = (e) => {
                e.preventDefault();
                localStorage.removeItem('username');
                window.location.reload();
            };
        } else {
            link.textContent = 'Login';
            link.href = '/pages/login.html';
            link.onclick = null;
        }
    });
}

// Run when the script loads
document.addEventListener('DOMContentLoaded', updateAuthState); 
