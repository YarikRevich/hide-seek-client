/*
Shader which implements basic torch lightning
*/

package shaders

var ScreenSize vec2

func Fragment(position vec4, texCoord vec2, color vec2) vec4 {
	coord := texCoord.xy / ScreenSize
	leftBottom := vec2(0.3)
	rightTop := vec2(0.7)

	m = max(rightTop-coord, coord-leftBottom)
	dist := length(max(vec2(0.0), m)) + min(0.0, max(m.x, m.y))
	return vec4(mix(vec3(.9, .4, 0.0), vec3(0.0, 0.0, 0.0), smoothstep(0.4, 0.20, dist)), 1.0)
}

// vec2 coord = (gl_FragCoord.xy / u_resolution);
// coord -= 0.5;
// coord *= vec2(u_resolution.x/u_resolution.y,1.0);
// coord += 0.5;

// vec3 color = vec3(0.0);
// color = vec3(rect(coord, vec2(0.7, 0.7)));
// color *= vec3(.9, .4, 0.0);

// vec2 tl = vec2(0.3);
// vec2 br = vec2(0.7);
// // float dx = max()
//     float dx = coord.x - .5;
//     float dy = coord.y - .5;
// float dist = length((dot(dy, dy)*3.14));
// vec2 d = max(tl- coord, coord - br);
// float dist = length(max(vec2(0.0, 0.0), d)) + min(0.0, max(d.x, d.y));
// vec3 c = mix( vec3(.9,.4,0.0), vec3(0.0, 0.0, 0.0), smoothstep(0.4, 0.20, dist) );
// gl_FragColor = vec4(c , 1.0);
