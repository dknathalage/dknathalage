export interface Employment {
    role: string;
    user: {
        uid: string;
        phone: string | null;
    };
    business: {
        uid: string;
        name: string;
    }
}

export interface User {
    uid: string;
    phone: string | null;
}

export interface Business {
    uid: string;
    name: string;
}