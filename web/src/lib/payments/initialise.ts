import { loadStripe } from "@stripe/stripe-js";
import type { Stripe } from "@stripe/stripe-js";

let stripePromise: Promise<Stripe | null>;

export const initialiseStripe = async () => {
    if (!stripePromise) {
        stripePromise = loadStripe("pk_test_51Gzu3YKjvXvpGllSwPBbMOkCzPDLWziHgX8J9WreUZncwNh4aFeCPEtjWKgmwqTBYVDWRFiH5jYXmsZQVbOZUVYr00b6uJNHk3");
    }

    return stripePromise;
}