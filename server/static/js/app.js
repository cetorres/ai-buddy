function setTheme (toggle = false) {
  const mode = 'auto';
  const userMode = localStorage.getItem('bs-theme');
  const sysMode = window.matchMedia('(prefers-color-scheme: light)').matches;
  let useSystem = mode === 'system' || (!userMode && mode === 'auto');
  let modeChosen = useSystem ? 'system' : mode === 'dark' || mode === 'light' ? mode : userMode;

  if (toggle) {
    if (modeChosen == 'light') modeChosen = 'dark';
    else modeChosen = 'light';
    useSystem = false;
  }

  if (useSystem) {
    localStorage.removeItem('bs-theme');
  } else {
    localStorage.setItem('bs-theme', modeChosen);
  }

  document.documentElement.setAttribute('data-bs-theme', useSystem ? (sysMode ? 'light' : 'dark') : modeChosen);
  document.getElementById('iconMoon').style.display = 'none';
  document.getElementById('iconSun').style.display = 'none';
  if (useSystem) {
    if (sysMode) document.getElementById('iconSun').style.display = 'block';
    else document.getElementById('iconMoon').style.display = 'block';
  }
  else {
    if (modeChosen == 'dark') document.getElementById('iconMoon').style.display = 'block';
    else document.getElementById('iconSun').style.display = 'block';
  }
}

setTheme();
document.getElementById('modeButton').addEventListener('click', () => setTheme(true));
window.matchMedia('(prefers-color-scheme: light)').addEventListener('change', () => setTheme());