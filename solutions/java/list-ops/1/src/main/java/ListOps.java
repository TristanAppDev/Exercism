import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.List;
import java.util.function.BiFunction;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;

class ListOps {

    static <T> List<T> append(List<T> list1, List<T> list2) {
        return Stream.of(list1, list2).flatMap(Collection::stream).collect(Collectors.toList());
    }

    static <T> List<T> concat(List<List<T>> listOfLists) {
        return listOfLists.stream().flatMap(List::stream).collect(Collectors.toList());
    }

    static <T> List<T> filter(List<T> list, Predicate<T> predicate) {
        return list.stream().filter(predicate).toList();
    }

    static <T> int size(List<T> list) {
        return list.size();
    }

    static <T, U> List<U> map(List<T> list, Function<T, U> transform) {
        return list.stream().map(transform).toList();
    }

    static <T> List<T> reverse(List<T> list) {
        return IntStream.range(0, list.size()).map(i -> list.size() - 1 - i).mapToObj(list::get).toList();
    }

    static <T, U> U foldLeft(final List<T> list, final U initial, final BiFunction<U, T, U> f) {
        if (list.isEmpty()) {
            return initial;
        }

        return foldLeft(new ArrayList<>(list.subList(1, list.size())), f.apply(initial, list.get(0)), f);
    }

    static <T, U> U foldRight(final List<T> list, final U initial, final BiFunction<T, U, U> f) {
        if (list.isEmpty()) {
            return initial;
        }

        return f.apply(list.get(0), foldRight(new ArrayList<>(list.subList(1, list.size())), initial, f));
    }

    private ListOps() {
        // No instances.
    }

}
