const controller = new ScrollMagic.Controller();

function headerScrollInit() {
	new ScrollMagic.Scene({
		triggerElement: "#scroll-trigger",
		triggerHook: 0
	}).setClassToggle(".header", "scrolled")
		//.addIndicators()
		.addTo(controller);
}

function heroParallaxInit() {
	let opacityTween = TweenMax.to(".parallax-mirror", 1, {opacity: 0.1, ease: Linear.easeNone});

	let opacityScene = new ScrollMagic.Scene({
		triggerElement: ".page-content",
		duration: 500
	}).setTween(opacityTween)
		.setPin(".parallax-mirror", {pushFollowers: false})
		//.addIndicators()
		.addTo(controller);
}

function changeTheme() {
	let body = $("body");
	if (body.hasClass("theme-light")) {
		body.toggleClass("theme-dark").toggleClass("theme-light");
	} else {
		body.toggleClass("theme-light").toggleClass("theme-dark");
	}
	$(this).toggleClass("dark").toggleClass("light");
}

function themeInit() {
	$(".theme").click(changeTheme);
}

function activePageInit(){
	$(`.main_nav a[href$='${window.location.pathname}']`)
		.parent()
		.toggleClass("active");
}

$(document).ready(function () {
	headerScrollInit();

	activePageInit();

	themeInit();

	//heroParallaxInit();
});