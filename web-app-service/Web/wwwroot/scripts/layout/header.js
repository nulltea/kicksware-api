const controller = new ScrollMagic.Controller();

function headerScrollInit() {
	new ScrollMagic.Scene({
		triggerElement: "#scroll-trigger",
		triggerHook: 0
	}).setClassToggle(".header", "scrolled")
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
	let baseRoute = window.location.pathname.split("/")[1];
	if (baseRoute){
		$(`.main_nav a[href*='/${window.location.pathname.split("/")[1]}/']`)
			.parent()
			.toggleClass("active");
	} else {
		$(`.main_nav a[href$='/']`).parent()
			.toggleClass("active");
	}
}

$(document).ready(function () {
	headerScrollInit();

	activePageInit();

	themeInit();

	//heroParallaxInit();
});