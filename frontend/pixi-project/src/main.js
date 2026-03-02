import { Application, Assets, Sprite } from "pixi.js";

(async () => {
  // Create a new application
  const app = new Application();

  // Initialize the application
  await app.init({ background: "#1099bb", resizeTo: window });

  // Append the application canvas to the document body
  document.getElementById("pixi-container").appendChild(app.canvas);

  // Load the bunny texture
  const texture = await Assets.load("/assets/bunny.png");

  // Create a bunny Sprite
  const bunny = new Sprite(texture);

  // Center the sprite's anchor point
  //bunny.anchor.set(0.5);

  // Move the sprite to the center of the screen
  bunny.position.set(0, app.screen.height / 2);

  bunny.scale.set(2);

  // Add the bunny to the stage
  app.stage.addChild(bunny);

  // Listen for animate update
  app.ticker.add((time) => {
    // Just for fun, let's rotate mr rabbit a little.
    // * Delta is 1 if running at 100% performance *
    // * Creates frame-independent transformation *
    //bunny.rotation += 0.01 * time.deltaTime;

    bunny.rotation += 0.1 * time.deltaTime;
    bunny.rotation -= 0.2 * time.deltaTime;
    bunny.rotation += 0.1 * time.deltaTime;
   
    

    bunny.x =  (bunny.x + 0.5 * time.deltaTime) % app.screen.width;


  });
})();

/*
const character = new PIXI.Container();

const body = PIXI.Sprite.from("body.png");
const arm = PIXI.Sprite.from("arm.png");

arm.anchor.set(0.1, 0.5);   // pivot at shoulder

character.addChild(body);
character.addChild(arm);

app.stage.addChild(character);






const frames = [
  PIXI.Texture.from("wave1.png"),
  PIXI.Texture.from("wave2.png"),
  PIXI.Texture.from("wave3.png"),
  PIXI.Texture.from("wave4.png")
];

const character = new PIXI.AnimatedSprite(frames);

character.animationSpeed = 0.2;
character.play();
*/


//SPRINTSHEET LEARN THE EXTENSIONS