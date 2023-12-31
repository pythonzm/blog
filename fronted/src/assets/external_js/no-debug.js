var minimalUserResponseInMiliseconds = 200
function check() {
  console.clear()
  var before = new Date().getTime()
  debugger; var after = new Date().getTime()
  if (after - before > minimalUserResponseInMiliseconds) {
    document.write(' Dont open Developer Tools. ')
    self.location.replace(window.location.protocol + window.location.href.substring(window.location.protocol.length))
  } else {
    before = null
    after = null
  }
  setTimeout(check, 100)
}

check()

window.onload = function() {
  document.addEventListener('contextmenu', function(e) {
    e.preventDefault()
  }, false)
  document.addEventListener('keydown', function(e) {
    if (e.ctrlKey && e.shiftKey && e.code === 73) {
      disabledEvent(e)
    }
    if (e.ctrlKey && e.shiftKey && e.code === 74) {
      disabledEvent(e)
    }
    if (e.code === 83 && (navigator.platform.match('Mac') ? e.metaKey : e.ctrlKey)) {
      disabledEvent(e)
    }
    if (e.ctrlKey && e.code === 85) {
      disabledEvent(e)
    }
    if (e.code === 123) {
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
}
