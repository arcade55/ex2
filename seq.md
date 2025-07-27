```mermaid

sequenceDiagram
    participant User
    participant Browser
    participant Datastar (Client-side)
    participant SortableJS (Client-side)

    Note over User,Browser: **Initial Page Load**
    Browser->>Datastar: **Load `datastar.min.js`** [3]
    activate Datastar
    Datastar->>Browser: **Initialise `$orderInfo` signal** with "Initial order" via `data-signals-order-info` [1].
    Datastar->>Browser: **Set up `data-text="$orderInfo"` binding** on `<p>` element [1].
    Datastar->>Browser: **Set up `data-on-reordered` listener** on `div#sortContainer` [1].
    deactivate Datastar

    Browser->>SortableJS: **Load SortableJS library** from CDN [1].
    activate SortableJS
    SortableJS->>Browser: **Initialise SortableJS** on `div#sortContainer` [1].
    Note over SortableJS: Configures `onEnd` callback to dispatch custom event [1].
    deactivate SortableJS

    User->>Browser: **Drags and drops an item** within `div#sortContainer`
    activate Browser
    Note over Browser,SortableJS: SortableJS library handles the drag-and-drop UI [1].
    SortableJS->>Browser: **Drag-and-drop action completes** (triggers `onEnd` callback) [1].
    activate SortableJS
    SortableJS-->>Browser: **Dispatches a `CustomEvent` named `reordered`** on `div#sortContainer` [1].
    Note over SortableJS: Event `detail` includes `orderInfo` (e.g., "Moved from 0 to 2") [1].
    deactivate SortableJS

    Browser->>Datastar: **`reordered` custom event received** by `data-on-reordered` listener [1].
    activate Datastar
    Datastar->>Datastar: **Evaluates the expression `$orderInfo = event.detail.orderInfo`** [1].
    Datastar->>Browser: **Updates the internal `$orderInfo` signal** with the new value from the event [2].
    Datastar-->>Browser: **Automatically re-renders** the `<p>` element due to `data-text="$orderInfo"` binding [1].
    deactivate Datastar
    deactivate Browser

    Note over User,Browser: **User sees the updated order information** displayed in the paragraph.