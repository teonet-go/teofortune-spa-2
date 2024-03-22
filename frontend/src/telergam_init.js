/*!
 * Telegram initialisation
 */
(() => {
    'use strict'
    let tg = window.Telegram;
    if (tg != undefined) {
        if (tg.WebApp != undefined && tg.WebApp.initData != undefined) {
            // let safe = tg.WebApp.initData;
            tg.WebApp.backgroundColor = '#3d3d3d';
            tg.WebApp.headerColor = '#212121';
            tg.WebApp.MainButton.hide();
            tg.WebApp.expand();

            tg.WebApp.onEvent('mainButtonClicked', function () {
                // When you click on the main button, we send data to bot in string
                tg.WebApp.sendData("some string that we need to send");
            });
        }
    }
})()
