import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import css from 'rollup-plugin-css-only';
import json from '@rollup/plugin-json';
import svelteSVG from "rollup-plugin-svelte-svg";
import replace from '@rollup/plugin-replace';

const production = !process.env.ROLLUP_WATCH;

export default {
	input: 'src/main.js',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'app',
		file: process.env.TWITCH == 'true' ? 'twitch-ext/build/bundle.js' : 'public/build/bundle.js'
	},
	plugins: [
		replace({
			preventAssignment: true,
      process: JSON.stringify({
        env: {
          isTwitchExt: process.env.TWITCH == 'true',
					apiBase: process.env.TWITCH_RIG != 'true' ? 'https://base64party.gomagames.com' : 'http://localhost:4000'
        }
      }),
    }),
		svelte({
			compilerOptions: {
				// enable run-time checks when not in production
				dev: !production
			}
		}),
		svelteSVG(),
		// we'll extract any component CSS out into
		// a separate file - better for performance
		css({ output: 'bundle.css' }),

		// If you have external dependencies installed from
		// npm, you'll most likely need these plugins. In
		// some cases you'll need additional configuration -
		// consult the documentation for details:
		// https://github.com/rollup/plugins/tree/master/packages/commonjs
		resolve({
			browser: true,
			dedupe: ['svelte']
		}),
		commonjs(),

		// Watch the `public` directory and refresh the
		// browser on changes when not in production
		!production && livereload(process.env.TWITCH == 'true' ? 'twitch-ext' : 'public'),

		// If we're building for production (npm run build
		// instead of npm run dev), minify
		production && terser(),

    // import .json
    json()
	],
	watch: {
		clearScreen: false
	}
};
