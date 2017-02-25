
const char *
get_kbd_state(Display *dpy)
{
    XkbStateRec state;
    const char *ru     = "Ru"; // if state group kbd = 1
    const char *en     = "En"; // if state group kbd = 0
    const char *no_kbd = "**"; // if keyboard any error

    if (XkbGetState(dpy, XkbUseCoreKbd, &state) == Success) {
        if (state.group == 1) {
                 return ru;
        } else {
                 return en;
        }
    }
    return no_kbd;

}
