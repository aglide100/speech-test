import { RefObject } from "react";
import { useDomEvent, MotionValue, animate } from "framer-motion";
import { mix } from "@popmotion/popcorn";
import { debounce } from "lodash";

// thanks for nice animation
// https://codesandbox.io/s/app-store-ui-using-react-and-framer-motion-1xdl9?file=/src/utils/use-wheel-scroll.ts

interface Constraints {
    top: number;
    bottom: number;
}

const deltaThreshold = 5;
const elasticFactor = 0.2;

function springTo(value: MotionValue, from: number, to: number) {
    if (value.isAnimating()) return;

    animate(value, to, {
        velocity: value.getVelocity(),
        stiffness: 400,
        damping: 40,
    });
}

const debouncedSpringTo = debounce(springTo, 100);

export function useWheelScroll(
    ref: RefObject<Element>,
    y: MotionValue<number>,
    constraints: Constraints | null,
    onWheelCallback: (e: WheelEvent) => void,
    isActive: boolean
) {
    const onWheel = (event: Event) => {
        event.preventDefault();

        const wheelEvent = event as WheelEvent;
        const currentY = y.get();
        let newY = currentY - wheelEvent.deltaY;
        let startedAnimation = false;
        const isWithinBounds =
            constraints &&
            newY >= constraints.top &&
            newY <= constraints.bottom;

        if (constraints && !isWithinBounds) {
            newY = mix(currentY, newY, elasticFactor);

            if (newY < constraints.top) {
                if (wheelEvent.deltaY <= deltaThreshold) {
                    springTo(y, newY, constraints.top);
                    startedAnimation = true;
                } else {
                    debouncedSpringTo(y, newY, constraints.top);
                }
            }

            if (newY > constraints.bottom) {
                if (wheelEvent.deltaY >= -deltaThreshold) {
                    springTo(y, newY, constraints.bottom);
                    startedAnimation = true;
                } else {
                    debouncedSpringTo(y, newY, constraints.bottom);
                }
            }
        }

        if (!startedAnimation) {
            y.stop();
            y.set(newY);
        } else {
            debouncedSpringTo.cancel();
        }

        onWheelCallback(wheelEvent);
    };

    useDomEvent(ref, "wheel", isActive ? onWheel : undefined, {
        passive: false,
    });
}
