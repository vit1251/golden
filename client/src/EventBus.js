
import { EventEmitter } from 'events';

import { store } from './Storage.js';


class EventBus extends EventEmitter {

    constructor() {
        super();
        this.active = false;
        this.socket = new WebSocket('ws://127.0.0.1:8080/api/v1');
        this.socket.addEventListener('open', (event) => {
            console.log(`Connect complete...`);
            this.active = true;
            if (this.reqQueue.length > 0) {
                console.log(`--- Queue requests ---`);
                for (const req of this.reqQueue) {
                    this.rawInvoke(req);
                }
                console.log(`--- Queue requests ---`);
            }
        });
        this.socket.addEventListener('message', (event) => {
            if (event.type === "message") {
                const {data = ''} = event;
                const msg = JSON.parse(data);
                console.log(msg);
                /* Common processing */
                this.emit('event', msg);
                /* Redux processing */
                store.dispatch(msg);
            } else {
                console.warn(`Wrong message.`);
            }
        });
        this.reqQueue = [];
    }

    rawInvoke(req) {
        console.log(`Use raw request`);
        const packet = JSON.stringify(req);
        this.socket.send(packet);
    }

    queueInvoke(req) {
        console.log(`Use queue request`);
        this.reqQueue.push(req);
    }

    /**
     * Invoke remote commands
     *
     * @param {Object} req
     */
    invoke(req) {
        console.log(`Invoke ${req}`);
        if (this.active) {
            this.rawInvoke(req);
        } else {
            this.queueInvoke(req);
        }
    }

}

export const eventBus = new EventBus();
