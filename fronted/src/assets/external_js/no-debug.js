window.onload = function() {
  document.addEventListener('contextmenu', function(e) {
    e.preventDefault()
  }, false)
  document.addEventListener('keydown', function(e) {
    if (e.ctrlKey && e.shiftKey && e.code === 'KeyI') {
      disabledEvent(e)
    }
    if (e.ctrlKey && e.shiftKey && e.code === 'KeyJ') {
      disabledEvent(e)
    }
    if (e.ctrlKey && e.shiftKey && e.code === 'KeyC') {
      disabledEvent(e)
    }
    if (e.code === 'KeyS' && (navigator.userAgent.match('/Macintosh/i') ? e.metaKey : e.ctrlKey)) {
      disabledEvent(e)
    }
    if (e.ctrlKey && e.code === 'KeyU') {
      disabledEvent(e)
    }
    if (e.code === 'F12') {
      disabledEvent(e)
    }
  }, false)
  function disabledEvent(e) {
    if (e.stopPropagation) {
      e.stopPropagation()
    } else if (window.Event) {
      window.Event.stopPropagation = true
    }
    e.preventDefault()
    return false
  }
  window.addEventListener('resize', function() {
    if ((window.outerHeight - window.innerHeight) > 100) {
      document.write('Bingo')
    }
  })
}
