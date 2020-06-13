function handAssetsInit(){
	const controller = new ScrollMagic.Controller();

	$(".info-section .trigger").each(function () {
		new ScrollMagic.Scene({
			triggerElement: this,
			// offset: 100
		}).setClassToggle($(this).find("~ .hand-asset")[0], "pushed")
			// .addIndicators()
			.addTo(controller);
	})
}

function creatorWindowInit(){
	const controller = new ScrollMagic.Controller();

	new ScrollMagic.Scene({
		triggerElement: "#bio-trigger",
		offset: isMobile() ? 100 : 200
	}).setClassToggle("#creator-window", "active")
		// .addIndicators()
		.addTo(controller);
}


$(document).ready(function () {
	"use strict";
	handAssetsInit();

	creatorWindowInit();
});