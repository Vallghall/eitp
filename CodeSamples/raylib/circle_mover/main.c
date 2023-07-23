#include "raylib.h"

#define WIDTH 500
#define HEIGHT 500
#define TITLE "Move this circle"
#define RADIUS 40.0f
#define DECOY 0.995f
#define SPEED_LIMIT 10.0f
#define limit(value) ((value > SPEED_LIMIT) ? SPEED_LIMIT : (value < -SPEED_LIMIT) ? -SPEED_LIMIT : value)

int main() {
    bool isHeld = false;
    Vector2 speed = {
        0.0f,
        0.0f,
    };
    Vector2 center = {
            (float) WIDTH/2,
            (float) HEIGHT/2,
    };

    InitWindow(WIDTH, HEIGHT, TITLE);
    SetTargetFPS(60);

    // GAME LOOP
    while (!WindowShouldClose()) {
        // UPDATE

        {
            // logic for holding and moving a circle with a mouse
            Vector2 cursorPos = GetMousePosition();
            if (CheckCollisionPointCircle(cursorPos, center, RADIUS) || isHeld) {
                if (IsMouseButtonPressed(MOUSE_BUTTON_LEFT)) {
                    isHeld = true;
                    speed.x = 0.0f;
                    speed.y = 0.0f;
                }

                if (IsMouseButtonDown(MOUSE_BUTTON_LEFT)) {
                    Vector2 delta = GetMouseDelta();

                    // Check collision with borders
                    {
                        float newX = center.x + delta.x;
                        float newY = center.y + delta.y;

                        if (newX > RADIUS && newX < WIDTH - RADIUS)
                            center.x = newX;
                        if (newY > RADIUS && newY < HEIGHT - RADIUS)
                            center.y = newY;
                    }
                }
            }

            if (IsMouseButtonReleased(MOUSE_BUTTON_LEFT) && isHeld) {
                isHeld = false;

                Vector2 delta = GetMouseDelta();
                speed.x = limit(0.8f * delta.x);
                speed.y = limit(0.8f * delta.y);
            }

            // logic for circle movement when it's free
            if (!isHeld) {
                // Check collision with borders
                {
                    if (center.x <= RADIUS || center.x >= WIDTH  - RADIUS) speed.x *= -1.0f;
                    if (center.y <= RADIUS || center.y >= HEIGHT - RADIUS) speed.y *= -1.0f;
                }

                center.x += speed.x;
                center.y += speed.y;

                speed.x *= DECOY;
                speed.y *= DECOY;
            }
        }
        // UPDATE

        // DRAW
        BeginDrawing();
        {
            ClearBackground(BLACK);
            DrawCircleV(center, 40.0f, RED);
        }
        EndDrawing();
        // DRAW
    }
    // GAME LOOP
    return 0;
}
