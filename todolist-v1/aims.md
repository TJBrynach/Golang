🛣️ TUI ToDo List App — Roadmap
🧹 Phase 1: Polish & Stability

    Goal: Make the app smooth, reliable, and user-friendly.

    ✅ Improve error handling by displaying messages in the UI (not just fmt.Println)

    ✅ Add input validation (e.g., prevent blank tasks or duplicates)

    ✅ Refactor hard-coded CSV indices into named constants DONE

    ✅ Normalize text input (e.g., trim spaces, lowercasing)

    🔄 Use []Task directly for populating the table instead of [][]string

    🧪 Start writing unit tests for:

        Task creation

        Task deletion

        CSV read/write
        
        check these are done and delet if not
✨ Phase 2: UI/UX Enhancements

    Goal: Make it more visually appealing and intuitive.

    🎨 Add color coding (e.g., completed tasks in green, errors in red)

    ✅ Display task status using emojis (✅, ❌, ⏳)

    🧾 Add a bottom bar with keyboard shortcut hints

    🔍 Implement a search/filter box (e.g., for "incomplete" or keyword search)

    ⌨️ Add keyboard shortcut to edit a task title

📦 Phase 3: Data Model Expansion

    Goal: Give users more control and depth.

    🗓️ Add optional due dates with visual reminders (e.g., red if overdue)

    📌 Add task priorities (low, medium, high)

    🗃️ Add categories or tags for task grouping

    🔁 Support recurring tasks (daily, weekly)

💾 Phase 4: Data Storage Upgrade

    Goal: Improve scalability and flexibility.

    🔄 Migrate from CSV to:

        JSON (for more complex task data)

        BoltDB (embedded key/value DB)

    ✍️ Implement graceful autosave (on every action, or debounce)

    🧹 Add a "purge completed tasks" option

🔐 Phase 5: Robustness & Portability

    Goal: Make it production-ready.

    🧪 Write full integration tests (end-to-end task flows)

    💻 Add support for cross-platform builds with a Makefile or shell script

    🗃️ Add a command-line flag to specify a different task file

    🧵 Improve concurrency safety if you add background features (e.g., reminders)

🌐 Phase 6: Going Beyond the Terminal

    Goal: Expand your app’s reach.

    ☁️ Sync to the cloud (GitHub Gist, Dropbox, or your own API)

    🌍 Add web version (React/Go back end)

    📱 Port to mobile using Gio or Flutter via Go bindings

    🔔 Add push notifications or system tray reminders

🏁 Final Vision

A fast, keyboard-driven task manager that's:

    Fully local-first and offline

    Sync-capable for cloud backups

    Clean UI with priority/status/dates

    Hackable and scriptable (maybe with a config file or plugin support)