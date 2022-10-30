<template>
	<div id="app">
    <div class="container">
      <StageArea></StageArea>
    </div>
<!--		<HelloWorld msg="Hello World Guark Vue Template" />-->
	</div>
</template>

<style lang="scss">
body, html {
	width: 100%;
	padding: 0;
	margin: 0;
  background-color: #1f2931;
}

#app {
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	height: 100%;
	width: 100%;
}

// Disable image drag and drop.
img {
	-webkit-user-drag: none;
	-khtml-user-drag: none;
	-moz-user-drag: none;
	-o-user-drag: none;
	user-drag: none;
}

* {
	user-select: none;
	-moz-user-select: none;
	-webkit-user-select: none;
	-ms-user-select: none;
	touch-callout: none;
	-webkit-touch-callout: none;
}

input {
	user-select: auto;
	-moz-user-select: auto;
	-webkit-user-select: auto;
	-ms-user-select: auto;
}

//.container{
//  padding: 20px;
//}

.test{
  height: 20px;
  background-color: #42b983;
}

</style>

<script>
import g from "guark"
import HelloWorld from '@/components/HelloWorld.vue'
import StageArea from "@/components/StageArea";

export default {
	name: "App",

	components: {
    StageArea,
		HelloWorld,
	},

	created()
	{
		g.env().then(env =>
		{
			// Disable right click menu on production, you can implement your own!
			if (!!env.dev_mode) {
				document.addEventListener('contextmenu', e => e.preventDefault());
			}
		})
	},
}
</script>