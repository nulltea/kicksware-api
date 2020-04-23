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

function handleHeroBrightness(imageSrc, callback) {
	let img = document.createElement("img");
	img.src = imageSrc;
	img.style.display = "none";
	document.body.appendChild(img);
	let colorSum = 0;
	img.onload = function() {
		let canvas = document.createElement("canvas");
		canvas.width = this.width;
		canvas.height = this.height;

		let ctx = canvas.getContext("2d");
		ctx.drawImage(this,0,0);

		let imageData = ctx.getImageData(0,0,canvas.width,canvas.height);
		let data = imageData.data;
		let r,g,b,avg;

		for(let x = 0, len = data.length; x < len; x+=4) {
			r = data[x];
			g = data[x+1];
			b = data[x+2];

			avg = Math.floor((r+g+b)/3);
			colorSum += avg;
		}

		let brightness = Math.floor(colorSum / (this.width*this.height));
		callback(brightness);
	}
}

$(document).ready(function () {
	headerScrollInit();

	activePageInit();

	themeInit();
});