'use strict'

/* https://embed.im/snow */
var embedimSnow = document.getElementById('embedim--snow')

if (!embedimSnow) {
  var embRand = function embRand(a, b) {
    return Math.floor(Math.random() * (b - a + 1)) + a
  }

  var embCSS = '.embedim-snow{position: absolute;width: 10px;height: 10px;background: white;border-radius: 50%;margin-top:-10px}'
  var embHTML = ''

  for (var i = 1; i < 200; i++) {
    embHTML += '<i class="embedim-snow"></i>'
    var rndX = embRand(0, 1000000) * 0.0001
    var rndO = embRand(-100000, 100000) * 0.0001
    var rndT = (embRand(3, 8) * 10).toFixed(2)
    var rndS = (embRand(0, 10000) * 0.0001).toFixed(2)
    embCSS += '.embedim-snow:nth-child(' + i + '){' + 'opacity:' + (embRand(1, 10000) * 0.0001).toFixed(2) + ';' + 'transform:translate(' + rndX.toFixed(2) + 'vw,-10px) scale(' + rndS + ');' + 'animation:fall-' + i + ' ' + embRand(10, 30) + 's -' + embRand(0, 30) + 's linear infinite' + '}' + '@keyframes fall-' + i + '{' + rndT + '%{' + 'transform:translate(' + (rndX + rndO).toFixed(2) + 'vw,' + rndT + 'vh) scale(' + rndS + ')' + '}' + 'to{' + 'transform:translate(' + (rndX + rndO / 2).toFixed(2) + 'vw, 105vh) scale(' + rndS + ')' + '}' + '}'
  }

  embedimSnow = document.createElement('div')
  embedimSnow.id = 'embedim--snow'
  embedimSnow.innerHTML = '<style>#embedim--snow{position:fixed;left:0;top:0;bottom:0;width:100vw;height:100vh;overflow:hidden;z-index:9999999;pointer-events:none}' + embCSS + '</style>' + embHTML
  document.body.appendChild(embedimSnow)
}
